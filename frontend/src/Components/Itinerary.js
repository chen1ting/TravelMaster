import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { validSessionGuard } from "../common/common";
import {
  ENDPOINT,
  getItinerary,
  getActivitiesByFilter,
  saveItineraryChanges,
} from "../api/api";
import { useNavigate } from "react-router-dom";
import StarRatings from "react-star-ratings";

import {
  Box,
  Heading,
  Switch,
  Text,
  FormLabel,
  Image,
  Button,
  Popover,
  Portal,
  PopoverContent,
  PopoverTrigger,
  Badge,
  Input,
} from "@chakra-ui/react";
import {
  AddIcon,
  ArrowLeftIcon,
  ArrowRightIcon,
  CloseIcon,
  InfoIcon,
} from "@chakra-ui/icons";

const EditItinerary = () => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(true);
  const [notifMsg, setNotifMsg] = useState("");
  const [notifMsg2, setNotifMsg2] = useState("");
  const [isError, setIsError] = useState(false);
  const [timeBins, setTimeBins] = useState([]);
  const [itineraryMap, setItineraryMap] = useState(new Map());

  const [timeBinsCopy, setTimeBinsCopy] = useState([]);
  const [itineraryMapCopy, setItineraryMapCopy] = useState(new Map());
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [curDate, setCurDate] = useState(new Date());
  const [itineraryResp, setItineraryResp] = useState(null);
  const session_token = window.sessionStorage.getItem("session_token");
  const [visualMode, setVisualMode] = useState(true);
  const [itiName, setItiName] = useState("");

  const { id } = useParams();
  useEffect(() => {
    validSessionGuard(navigate, "/");
    getItinerary(
      parseInt(id),
      session_token,
      setIsLoading,
      setNotifMsg,
      setTimeBins,
      setItineraryMap,
      setStartDate,
      setEndDate,
      setItineraryResp,
      setCurDate,
      setTimeBinsCopy,
      setItineraryMapCopy,
      setItiName
    );
  }, [id, navigate, session_token]);

  const getStartIdx = () =>
    (curDate.getTime() - startDate.getTime()) / (1000 * 60 * 60);

  const saveChanges = async () => {
    const saved = await saveItineraryChanges(
      parseInt(id),
      timeBins,
      window.sessionStorage.getItem("session_token"),
      startDate.getTime() / 1000,
      itiName
    );
    if (saved == null) {
      setNotifMsg2("Something went wrong saving the new itinerary");
      setIsError(true);
    } else {
      setNotifMsg2("Successfully saved itinerary");
    }

    await new Promise((resolve) => setTimeout(resolve, 3000));
    setNotifMsg2("");
    setIsError(false);
  };

  return (
    <Box>
      {isLoading ? (
        <Text>Loading ... </Text>
      ) : notifMsg ? (
        <Text>{notifMsg}</Text>
      ) : (
        <Box
          w="100vw"
          display="flex"
          flexDir="column"
          alignItems="center"
          mb="16"
        >
          <Box mt="8" textAlign="center">
            <Heading size="md">Your generated itinerary from</Heading>
            <Heading size="md" my="3">
              <span>
                {new Date(itineraryResp.start_time * 1000).toLocaleString()}
              </span>
              <span>{" to "}</span>
              <span>
                {new Date(itineraryResp.end_time * 1000).toLocaleString()}
              </span>
              .
            </Heading>
          </Box>
          {notifMsg2 && (
            <Box
              px="5"
              py="3"
              color={'white'}
              bgColor={isError ? "orange.500" : "green.400"}
              borderRadius="15px"
            >
              {notifMsg2}
            </Box>
          )}

          <Box
            w="70%"
            display="flex"
            justifyContent="center"
            columnGap="6"
            alignItems="center"
            my="5"
          >
            <Button
              colorScheme="yellow"
              variant="outline"
              onClick={saveChanges}
            >
              Save changes made
            </Button>
            <Button
              colorScheme="teal"
              variant="outline"
              onClick={() => {
                setTimeBins([...timeBinsCopy]);
                setItineraryMap(new Map(itineraryMapCopy));
              }}
            >
              Reset changes made
            </Button>
            <FormLabel htmlFor="toggle-visual-mode" mb="0">
              Visual Mode: {visualMode ? "on" : "off"}
            </FormLabel>
            <Switch
              ml="-6"
              id="visual-mode"
              colorScheme="teal"
              isChecked={visualMode}
              onChange={() => setVisualMode((prev) => !prev)}
            />
          </Box>
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            columnGap="3"
          >
            <Text>Itinerary name:</Text>
            <Input
              type="text"
              value={itiName}
              onChange={(e) => setItiName(e.target.value)}
              w="200px"
            />
          </Box>
          <Text size="sm">
            (Tip: Change this to something you can remember it by!)
          </Text>
          <Box
            mt="8"
            mb="5"
            display="flex"
            justifyContent="center"
            alignItems="center"
            columnGap="16"
          >
            <ArrowLeftIcon
              w={5}
              h={5}
              cursor={
                curDate.getTime() === startDate.getTime()
                  ? "not-allowed"
                  : "pointer"
              }
              _hover={{
                color:
                  curDate.getTime() === startDate.getTime()
                    ? "gray.600"
                    : "blue.300",
              }}
              onClick={() => {
                if (curDate.getTime() === startDate.getTime()) {
                  return;
                }
                setCurDate(
                  (prev) =>
                    new Date(
                      prev.getFullYear(),
                      prev.getMonth(),
                      prev.getDate() - 1
                    )
                );
              }}
              color={
                curDate.getTime() === startDate.getTime() ? "gray.300" : "black"
              }
            />
            <Text fontSize="18" fontWeight="600" style={{ userSelect: "none" }}>
              {curDate.toLocaleDateString()}
            </Text>
            <ArrowRightIcon
              w={5}
              h={5}
              cursor={
                curDate.getTime() === endDate.getTime()
                  ? "not-allowed"
                  : "pointer"
              }
              _hover={{
                color:
                  curDate.getTime() === endDate.getTime() ? "gray.600" : "blue.300",
              }}
              onClick={() => {
                if (curDate.getTime() === endDate.getTime()) {
                  return;
                }
                setCurDate(
                  (prev) =>
                    new Date(
                      prev.getFullYear(),
                      prev.getMonth(),
                      prev.getDate() + 1
                    )
                );
              }}
              color={curDate.getTime() === endDate.getTime() ? "gray.300" : "black"}
            />
          </Box>

          <Box my="6" py="6" pr="8" bgColor="gray.200">
            {timeBins
              .slice(getStartIdx(), getStartIdx() + 24)
              .map((activity, index) =>
                TimeCell(
                  activity,
                  index,
                  visualMode,
                  navigate,
                  (activity) =>
                    setTimeBins((prev) =>
                      prev.map((_, i) =>
                        i === getStartIdx() + index ? activity : prev[i]
                      )
                    ),
                  [
                    {
                      day: curDate.getDay(),
                      start_time_offset: index,
                      end_time_offset: index + 1,
                    },
                  ]
                )
              )}
          </Box>
        </Box>
      )}
    </Box>
  );
};

const TimeCell = (
  activity,
  index,
  visualMode,
  navigate,
  replaceSelf,
  times
) => {
  const hour = index.toString().padStart(2, "0");
  const nxtHr = (index + 1).toString().padStart(2, "0");
  const label = `${hour}:00 - ${nxtHr}:00`;
  return (
    <Box
      display="flex"
      justifyContent="flex-start"
      alignItems="center"
      borderTop="0"
      borderBottom="0"
    >
      <Box
        minW="150px"
        display="flex"
        justifyContent="center"
        alignItems="flex-start"
      >
        {label}
      </Box>

      <PopoverWrapper
        activity={activity}
        navigate={navigate}
        replaceSelf={replaceSelf}
        times={times}
        label={label}
      >
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          w="100%"
          minW="350px"
          border="1px solid white"
        >
          {activity ? (
            <Box
              h="120px"
              w="full"
              display="flex"
              flexDir="column"
              alignItems="center"
              justifyContent="center"
            >
              {visualMode ? (
                <Image
                  w="full"
                  h="120px"
                  filter="auto"
                  brightness="90%"
                  objectFit="cover"
                  src={`${ENDPOINT}/activity-images/` + activity.image_names[0]}
                  alt={activity.name}
                  cursor="pointer"
                  _hover={{
                    brightness: "110%",
                  }}
                />
              ) : (
                <Box
                  filter="auto"
                  brightness="90%"
                  cursor="pointer"
                  _hover={{
                    brightness: "120%",
                  }}
                >
                  <Text>{activity.name}</Text>
                  <Text noOfLines="3">{activity.description}</Text>
                </Box>
              )}
            </Box>
          ) : (
            <Box
              h="30px"
              display="flex"
              justifyContent="center"
              alignItems="center"
              filter="auto"
              brightness="90%"
              cursor="pointer"
              w="full"
              _hover={{
                brightness: "120%",
                backgroundColor: "green.300",
              }}
            >
              <Text>Nothing here</Text>
            </Box>
          )}
        </Box>
      </PopoverWrapper>
    </Box>
  );
};

const PopoverWrapper = ({
  children,
  activity,
  navigate,
  replaceSelf,
  times,
  label,
}) => {
  return (
    <Popover placement="right" trigger="hover">
      <Box>
        <PopoverTrigger>{children}</PopoverTrigger>

        <Portal>
          <PopoverContent>
            {activity ? (
              <Box py="4" px="2">
                <Image
                  src={`${ENDPOINT}/activity-images/` + activity.image_names[0]}
                  alt={activity.name}
                  mb="4"
                />
                <Heading size="md">{activity.name}</Heading>
                <Box
                  my="2"
                  display="flex"
                  alignItems="center"
                  columnGap="2"
                  flexWrap="wrap"
                  rowGap="2"
                >
                  {activity.categories && activity.categories.map((cat) => (
                    <Badge key={cat} variant="outline" colorScheme="green">
                      {cat}
                    </Badge>
                  ))}
                </Box>

                <Box display="flex" alignItems="center" columnGap="3">
                  <StarRatings
                    rating={activity.average_rating}
                    starRatedColor="#F6E05E"
                    starDimension="14px"
                    numberOfStars={5}
                    name="rating"
                    starSpacing="2px"
                  />
                  <Text mt="1px">({activity.average_rating})</Text>
                </Box>

                <Text noOfLines="3" mt="3">
                  {activity.description}
                </Text>

                <Box
                  display="flex"
                  justifyContent="center"
                  alignItems="center"
                  columnGap="5"
                  mt="5"
                  mb="1"
                >
                  <Button
                    colorScheme="teal"
                    leftIcon={<InfoIcon />}
                    onClick={() => navigate(`/activity/${activity.id}`)}
                  >
                    More Info
                  </Button>
                  <Button
                    colorScheme="red"
                    leftIcon={<CloseIcon />}
                    onClick={() => replaceSelf(null)}
                  >
                    Remove
                  </Button>
                </Box>
              </Box>
            ) : (
              <SmallSearchActivity
                label={label}
                times={times}
                navigate={navigate}
                replaceSelf={replaceSelf}
              />
            )}
          </PopoverContent>
        </Portal>
      </Box>
    </Popover>
  );
};

const SmallSearchActivity = ({ label, times, navigate, replaceSelf }) => {
  const [activities, setActivities] = useState([]);
  const [pageNum, setPageNum] = useState(1);
  const pageSize = 3;
  const session_token = window.sessionStorage.getItem("session_token");

  return (
    <Box py="4" px="2">
      <Heading size="sm" mb="2">
        Add a new activity for {label}
      </Heading>
      <Input
        my="2"
        type="text"
        size="sm"
        placeholder="Search for an activity"
        onChange={(e) =>
          getActivitiesByFilter(
            e.target.value,
            pageNum,
            pageSize,
            times,
            session_token,
            setActivities
          )
        }
      />
      {activities.length > 0 && (
        <Box
          display="flex"
          justifyContent="flex-end"
          alignItems="center"
          columnGap="3"
        >
          <Text>Page: </Text>
          <ArrowLeftIcon
            w={3}
            h={3}
            cursor={pageNum === 1 ? "not-allowed" : "pointer"}
            _hover={{
              color: pageNum === 1 ? "gray" : "blue.300",
            }}
            onClick={() => {
              if (pageNum === 1) {
                return;
              }
              setPageNum((prev) => prev - 1);
            }}
          />
          <Text>{pageNum}</Text>
          <ArrowRightIcon
            w={3}
            h={3}
            cursor={activities.length < pageSize ? "not-allowed" : "pointer"}
            _hover={{
              color: activities.length < pageSize ? "gray" : "blue.300",
            }}
            onClick={() => {
              if (activities.length < pageSize) {
                return;
              }
              setPageNum((prev) => prev + 1);
            }}
          />
        </Box>
      )}
      <Box mb="6">
        {activities && activities.map((act) => (
          <Box my="3" borderBottom="1px solid white">
            <Image
              mt="4"
              h="150px"
              w="full"
              src={`${ENDPOINT}/activity-images` + act.image_names[0]}
              alt={act.name}
            />
            <Heading mt="2" size="sm">
              {act.name}
            </Heading>
            <Box my="2" display="flex" alignItems="center" columnGap="2">
              {act.categories && act.categories.map((cat) => (
                <Badge variant="outline" colorScheme="green">
                  {cat}
                </Badge>
              ))}
            </Box>
            <Box display="flex" alignItems="center" columnGap="3">
              <StarRatings
                rating={act.average_rating}
                starRatedColor="#F6E05E"
                starDimension="14px"
                numberOfStars={5}
                name="rating"
                starSpacing="2px"
              />
              <Text mt="1px">({act.average_rating})</Text>
            </Box>
            <Text noOfLines="3" my="2">
              {act.description}
            </Text>
            <Box
              display="flex"
              justifyContent="center"
              alignItems="center"
              columnGap="5"
              mt="5"
              mb="6"
            >
              <Button
                colorScheme="teal"
                leftIcon={<InfoIcon />}
                onClick={() => navigate(`/activity/${act.id}`)}
              >
                More Info
              </Button>
              <Button
                colorScheme="green"
                variant="outline"
                leftIcon={<AddIcon />}
                onClick={() => replaceSelf(act)}
              >
                Add
              </Button>
            </Box>
          </Box>
        ))}
      </Box>
    </Box>
  );
};

export default EditItinerary;
