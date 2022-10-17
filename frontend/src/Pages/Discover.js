import { useNavigate } from "react-router-dom";
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";
import { useEffect, useState } from "react";
import { ENDPOINT } from "../api/api";
import {
  FormControl,
  FormLabel,
  Button,
  Input,
  Box,
  Text,
  InputGroup,
  InputLeftAddon,
  Heading,
  Image,
  useDisclosure,
  Badge,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalCloseButton,
  ModalBody,
  ModalFooter,
  Checkbox,
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
  Textarea,
} from "@chakra-ui/react";
import StarRatings from "react-star-ratings";
import { ArrowLeftIcon, ArrowRightIcon } from "@chakra-ui/icons";
import { getActivitiesByFilter, sendCreateActivityReq } from "../api/api";
import { SearchIcon } from "@chakra-ui/icons";
import AsyncSelect from "react-select/async";
import Geocode from "react-geocode";
import { apiKey } from "../common/common";

Geocode.setApiKey(apiKey);
Geocode.setRegion("sg");

const Discover = () => {
  const navigate = useNavigate();
  const [searchText, setSearchText] = useState("");
  const [pageNum, setPageNum] = useState(1);
  const [activities, setActivities] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);
  const pageSize = 5;

  const { isOpen, onOpen, onClose } = useDisclosure();

  const fetchActivities = async () => {
    setIsLoading(true);
    await getActivitiesByFilter(
      searchText,
      pageNum,
      pageSize,
      [],
      window.sessionStorage.getItem("session_token"),
      setActivities
    );

    setIsLoading(false);
  };

  useEffect(() => {
    fetchActivities();
  }, [pageNum]);

  return (
    <Box w="100%">
      {isLoading ? (
        <Text>Loading...</Text>
      ) : (
        <Box w="100%" mb="32" mt="3">
          <Box
            pt="7"
            display="flex"
            justifyContent="flex-start"
            ml="10%"
            // bgColor="blue.50s0"
          >
            <Box>
              <Heading size="3xl">Discover</Heading>
              <Text>Your next adventure is waiting for you.</Text>
            </Box>
          </Box>
          {notifMsg && (
            <Box justifyContent="center" display="flex">
              <Box
                textAlign="center"
                w="50%"
                mt="3"
                bgColor={isError ? "red.400" : "green.500"}
                py="2"
                px="6"
                borderRadius="15px"
              >
                <Text fontSize="18">{notifMsg}</Text>
              </Box>
            </Box>
          )}
          <Box w="100%" mt="10" mb="8">
            <InputGroup
              size="lg"
              display="flex"
              alignItems="center"
              justifyContent="center"
            >
              <InputLeftAddon children={<SearchIcon />} />
              <Input
                placeholder="Search for an activity"
                w="50%"
                onChange={(e) => setSearchText(e.target.value)}
              />
              <Button ml="9" colorScheme="green" onClick={fetchActivities}>
                Search
              </Button>
            </InputGroup>
          </Box>

          <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
            mt="10"
          >
            <Box ml="20%">
              <Button onClick={onOpen} colorScheme="blue">
                Create an activity
              </Button>
            </Box>
            <Box
              minW="200px"
              display="flex"
              mr="20%"
              alignItems="center"
              columnGap="3"
            >
              <Text fontSize="23">Page: </Text>
              <ArrowLeftIcon
                w={4}
                h={4}
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
              <Text fontSize="23">{pageNum}</Text>
              <ArrowRightIcon
                w={4}
                h={4}
                cursor={
                  activities.length < pageSize ? "not-allowed" : "pointer"
                }
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
          </Box>

          <Box
            display="flex"
            flexDir="column"
            justifyContent="center"
            alignItems="center"
            mt="5"
          >
            {activities &&
              activities.map((act) => (
                <ActivityCard act={act} navigate={navigate} />
              ))}
          </Box>
        </Box>
      )}

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <CreateForm
          onClose={onClose}
          setNotifMsg={setNotifMsg}
          setIsError={setIsError}
          navigate={navigate}
        />
      </Modal>
    </Box>
  );
};

const ActivityCard = ({ act, navigate }) => {
  return (
    <Box
      w="1200px"
      display="flex"
      justifyContent="center"
      onClick={() => {
        navigate(`/activity/${act.id}`);
      }}
    >
      <Box
        cursor="pointer"
        filter="auto"
        brightness="90%"
        _hover={{
          brightness: "120%",
          backgroundColor: "teal.300",
        }}
        border="1px solid white"
        borderRadius="15px"
        display="flex"
        justifyContent="space-around"
        my="3"
      >
        <Box padding="5">
          <Image
            borderRadius="30px"
            w="400px"
            h="300px"
            src={`${ENDPOINT}/activity-images/${act.image_names[0]}`}
            alt={act.name}
          />
        </Box>
        <Box
          p="5"
          display="flex"
          flexDir="column"
          rowGap="2"
          ml="16"
          minW="400px"
          maxW="300px"
          flexWrap="wrap"
        >
          <Heading mt="2" fontSize="25">
            {act.name}
          </Heading>
          <Box
            my="2"
            display="flex"
            alignItems="center"
            columnGap="2"
            flexWrap="wrap"
            rowGap="2"
          >
            {act.categories &&
              act.categories.map((cat) => (
                <Badge variant="outline" colorScheme="green">
                  {cat}
                </Badge>
              ))}
          </Box>

          <Box display="flex" alignItems="center" columnGap="3">
            <StarRatings
              rating={act.average_rating}
              starRatedColor="#F6E05E"
              starDimension="20px"
              numberOfStars={5}
              name="rating"
              starSpacing="3px"
            />
            <Text mt="1px">({act.average_rating})</Text>
          </Box>
          <Text noOfLines="5" mt="8" fontSize="18">
            {act.description}
          </Text>
        </Box>
      </Box>
    </Box>
  );
};

const CreateForm = ({ onClose, setNotifMsg, setIsError, navigate }) => {
  const catsList = [
    "Museums",
    "Nature",
    "Hiking",
    "Educational",
    "Animals",
    "Food and beverage",
    "Special Events",
    "Concerts",
    "Live Music",
    "Local Culture",
    "Hidden Gems",
    "Luxurious",
    "Thrifty",
    "Shopping",
  ];
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [cats, setCats] = useState(new Set());
  const [picture, setPicture] = useState(null);
  const [isPaid, setIsPaid] = useState(false);
  const [loc, setLoc] = useState({});

  const daysList = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ];
  const [hours, setHours] = useState([
    7, 7, 7, 7, 7, 7, 7, 22, 22, 22, 22, 22, 22, 22,
  ]); // day: i, open: i, close i+7

  const [isOpen, setIsOpen] = useState([
    true,
    true,
    true,
    true,
    true,
    true,
    true,
  ]);

  const submitCreateActivityForm = async () => {
    onClose();
    const uid = window.sessionStorage.getItem("uid");
    if (!uid) {
      setNotifMsg("Create an account first to contribute an activity!");
      setIsError(true);
      await new Promise((resolve) => setTimeout(resolve, 3000));
      setNotifMsg("");
      setIsError(false);
      return;
    }
    const tmpHours = [...hours];
    for (let i = 0; i < 7; i++) {
      if (!isOpen[i]) {
        tmpHours[i] = 25;
        tmpHours[i + 7] = 25;
      }
    }

    const data = await sendCreateActivityReq(
      uid,
      title,
      isPaid,
      cats,
      desc,
      picture,
      tmpHours,
      loc
    );

    if (data.error) {
      setNotifMsg(data.error);
      setIsError(true);
      await new Promise((resolve) => setTimeout(resolve, 3000));
      setNotifMsg("");
      setIsError(false);
    } else {
      setNotifMsg("Successfully created activity");
      setIsError(false);
      await new Promise((resolve) => setTimeout(resolve, 800));
      navigate(`/activity/${data.activity_id}`);
    }
  };

  const loadOptions = (inputValue, callback) =>
    Geocode.fromAddress(inputValue).then(
      (response) => {
        callback(
          response.results.map((res) => ({
            value: res.geometry.location,
            label: res.formatted_address,
          }))
        );
      },
      (error) => {
        console.error(error);
      }
    );

  return (
    <ModalContent minW="600px">
      <ModalHeader>Create an activity</ModalHeader>
      <ModalCloseButton />
      <ModalBody pb={6}>
        <FormControl isRequired>
          <FormLabel>Activity Title</FormLabel>
          <Input
            placeholder="Enter a title for the activity"
            onChange={(e) => setTitle(e.target.value)}
          />
        </FormControl>

        <FormControl isRequired mt={4}>
          <FormLabel>Activity Description</FormLabel>
          <Textarea
            placeholder="Enter a brief description of the activity"
            onChange={(e) => setDesc(e.target.value)}
          />
        </FormControl>

        <FormControl isRequired my="5">
          <FormLabel>Upload an image for the activity</FormLabel>
          <input
            type="file"
            //style={{ display: 'none' }}
            onChange={(e) => {
              setPicture(e.target.files[0]);
            }}
          />
        </FormControl>

        <FormControl isRequired mt={4}>
          <FormLabel>Does this event require the users to pay?</FormLabel>
          <Checkbox
            size="md"
            colorScheme="green"
            onChange={(e) => setIsPaid(e.target.checked)}
          >
            Yes, this is a paid activity.
          </Checkbox>
        </FormControl>

        <FormControl isRequired mt={4}>
          <FormLabel>Select a few tags the activity falls under:</FormLabel>
          {catsList.map((cat) => (
            <Button
              m="1"
              variant="outline"
              borderRadius="15px"
              colorScheme={cats.has(cat) ? "whatsapp" : "facebook"}
              key={cat}
              onClick={() =>
                setCats((prev) =>
                  prev.has(cat)
                    ? new Set([...prev].filter((x) => x !== cat))
                    : new Set([...prev, cat])
                )
              }
            >
              {cat}
            </Button>
          ))}
        </FormControl>

        <FormControl isRequired mt={4}>
          <FormLabel>Location</FormLabel>
          <Box color="black">
            <AsyncSelect
              onChange={(opt) => {
                setLoc(opt.value);
              }}
              loadOptions={loadOptions}
            />
          </Box>
        </FormControl>

        <FormControl isRequired mt={4}>
          <FormLabel>Operating hours</FormLabel>
          <TableContainer>
            <Table size="sm">
              <Thead>
                <Tr>
                  <Th>Day</Th>
                  <Th>Is Open</Th>
                  <Th>Opening Time</Th>
                  <Th>Closing Time</Th>
                </Tr>
              </Thead>
              <Tbody>
                {daysList.map((day, i) => {
                  const open = new Date();
                  open.setHours(hours[i]);
                  open.setMinutes(0);

                  const close = new Date();
                  close.setHours(hours[i + 7]);
                  close.setMinutes(0);
                  return (
                    <Tr>
                      <Td>{day}</Td>
                      <Td>
                        <Checkbox
                          colorScheme="green"
                          isChecked={isOpen[i]}
                          onChange={(e) =>
                            setIsOpen((prev) =>
                              prev.map((b, idx) =>
                                idx === i ? e.target.checked : b
                              )
                            )
                          }
                        />
                      </Td>
                      {isOpen[i] ? (
                        <>
                          <Td>
                            <DatePicker
                              selected={open}
                              onChange={(date) =>
                                setHours((prev) =>
                                  prev.map((e, idx) =>
                                    idx === i ? date.getHours() : e
                                  )
                                )
                              }
                              showTimeSelect
                              showTimeSelectOnly
                              timeIntervals={60}
                              timeCaption="Time"
                              dateFormat="h:mm aa"
                            />
                          </Td>
                          <Td>
                            <DatePicker
                              selected={close}
                              onChange={(date) =>
                                setHours((prev) =>
                                  prev.map((e, idx) =>
                                    idx === i + 7 ? date.getHours() : e
                                  )
                                )
                              }
                              showTimeSelect
                              showTimeSelectOnly
                              timeIntervals={60}
                              timeCaption="Time"
                              dateFormat="h:mm aa"
                            />
                          </Td>
                        </>
                      ) : (
                        <>
                          <Td>-</Td>
                          <Td>-</Td>
                        </>
                      )}
                    </Tr>
                  );
                })}
              </Tbody>
            </Table>
          </TableContainer>
        </FormControl>
      </ModalBody>

      <ModalFooter>
        <Button colorScheme="blue" mr={3} onClick={submitCreateActivityForm}>
          Save
        </Button>
        <Button onClick={onClose}>Cancel</Button>
      </ModalFooter>
    </ModalContent>
  );
};

export default Discover;

export { ActivityCard };
