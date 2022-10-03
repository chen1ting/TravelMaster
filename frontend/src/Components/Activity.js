import {
  Box,
  Heading,
  Image,
  Text,
  Badge,
  Table,
  Thead,
  Tbody,
  Tfoot,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
  Textarea,
  Avatar,
  Input,
  Button,
  Divider,
  IconButton,
} from "@chakra-ui/react";
import { CheckIcon, CloseIcon, DeleteIcon, EditIcon } from "@chakra-ui/icons";
import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import {
  getActivityById,
  addReview,
  fetchUserInfo,
  sendUpdateReview,
} from "../api/api";
import StarRatings from "react-star-ratings";
import {
  withScriptjs,
  withGoogleMap,
  GoogleMap,
  Marker,
} from "react-google-maps";
import { apiKey } from "../common/common";

const Activity = () => {
  const { id } = useParams();
  const [isLoading, setIsLoading] = useState(true);
  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);
  const [act, setActivity] = useState(null);

  const [clicks, setClicks] = React.useState([]);
  const [zoom, setZoom] = React.useState(3); // initial zoom
  const [center, setCenter] = React.useState({
    lat: 0,
    lng: 0,
  });
  const onIdle = (m) => {
    console.log("onIdle");
    setZoom(m.getZoom());
    setCenter(m.getCenter().toJSON());
  };

  const onClick = (e) => {
    // avoid directly mutating state
    console.log(e.latLng);
    setClicks([...clicks, e.latLng]);
  };

  useEffect(() => {
    getActivityById(id, setActivity, setIsLoading);
  }, []);

  const render = (status) => {
    return <h1>{status}</h1>;
  };

  const daysList = ["sun", "mon", "tue", "wed", "thur", "fri", "sat"];

  return (
    <Box display="flex" justifyContent="center">
      {isLoading ? (
        <Text>Loading...</Text>
      ) : (
        <Box maxW="80%">
          <Box
            display="flex"
            justifyContent="center"
            mt="12"
            py="10"
            columnGap="32"
          >
            <Box py="8">
              <Image
                maxW="400px"
                h="350px"
                src={`http://localhost:8080/activity-images/${act.image_names[0]}`}
                alt={act.title}
                borderRadius="20px"
              />
              <Box
                mt="8"
                mb="2"
                display="flex"
                alignItems="center"
                columnGap="2"
              >
                {act.category.map((cat) => (
                  <Badge variant="outline" colorScheme="green">
                    {cat}
                  </Badge>
                ))}
              </Box>
              <Badge
                my="4"
                size="lg"
                colorScheme={act.paid ? "yellow" : "teal"}
              >
                {act.paid ? "PAID" : "FREE"}
              </Badge>
              <Box display="flex" alignItems="center" columnGap="3">
                <StarRatings
                  rating={act.rating_score}
                  starRatedColor="#F6E05E"
                  starDimension="20px"
                  numberOfStars={5}
                  name="rating"
                  starSpacing="3px"
                />
                <Text mt="2px">{act.review_counts} review(s)</Text>
              </Box>
            </Box>
            <Box maxW="700px">
              <Heading>{act.title}</Heading>
              <Text my="7">{act.description}</Text>
              <TableContainer>
                <Table size="md">
                  <Thead>
                    <Tr>
                      <Th>Day</Th>
                      <Th>Opening Time</Th>
                      <Th>Closing Time</Th>
                    </Tr>
                  </Thead>
                  <Tbody>
                    {daysList.map((day) => {
                      const open = new Date();
                      open.setHours(act[`${day}_opening_time`]);
                      open.setMinutes(0);

                      const close = new Date();
                      close.setHours(act[`${day}_closing_time`]);
                      close.setMinutes(0);
                      return (
                        <Tr>
                          <Td>{day}</Td>
                          <Td>
                            {open.getHours().toString().padStart(2, "0")}:
                            {close.getMinutes().toString().padStart(2, "0")}
                          </Td>
                          <Td>
                            {close.getHours().toString().padStart(2, "0")}:
                            {close.getMinutes().toString().padStart(2, "0")}
                          </Td>
                        </Tr>
                      );
                    })}
                  </Tbody>
                </Table>
              </TableContainer>
            </Box>
          </Box>

          <Box display="flex" justifyContent="center" my="12">
            <MyMapComponent
              isMarkerShown
              googleMapURL={`https://maps.googleapis.com/maps/api/js?key=${apiKey}&v=3.exp&libraries=geometry,drawing,places`}
              loadingElement={<div style={{ height: `100%` }} />}
              containerElement={
                <div style={{ height: `500px`, width: "700px" }} />
              }
              mapElement={<div style={{ height: `100%` }} />}
              latLng={{ lat: act.latitude, lng: act.longitude }}
            />
          </Box>

          <Divider mt="12" mb="16" />
          <Reviews
            reviews={act.review_list}
            aid={act.activity_id}
            setActivity={setActivity}
          />
        </Box>
      )}
    </Box>
  );
};

const Reviews = ({ reviews, aid, setActivity }) => {
  const [stars, setStars] = useState(0);
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);

  const submitReview = async () => {
    const code = await addReview(
      window.sessionStorage.getItem("session_token"),
      aid,
      title,
      desc,
      stars,
      setActivity
    );
    if (code === 405) {
      setIsError(true);
      setNotifMsg("You have already uploaded your review before.");
      await new Promise((resolve) => setTimeout(resolve, 3000));
      setIsError(false);
      setNotifMsg("");
      return;
    } else if (code !== 201) {
      setIsError(true);
      setNotifMsg("Something went wrong...");
      await new Promise((resolve) => setTimeout(resolve, 3000));
      setIsError(false);
      setNotifMsg("");
      return;
    }

    setIsError(false);
    setNotifMsg("Successfully uploaded your review.");
    setStars(0);
    setTitle("");
    setDesc("");
    await new Promise((resolve) => setTimeout(resolve, 3000));
    setNotifMsg("");
  };
  return (
    <Box mb="32">
      <Heading mt="3" mb="5">
        Reviews
      </Heading>
      {notifMsg && (
        <Text
          m="4"
          w="500px"
          px="5"
          py="3"
          bgColor={isError ? "orange.400" : "green.400"}
          borderRadius="15px"
        >
          {notifMsg}
        </Text>
      )}

      <Box display="flex" justifyContent="flex-start">
        <Avatar
          mt="1"
          mr="4"
          ml="2"
          src={`http://localhost:8080/avatars/${window.sessionStorage.getItem(
            "avatar_file_name"
          )}`}
        />
        <Box display="flex" flexDir="column" rowGap="5">
          <Box display="flex" columnGap="10" alignItems="center">
            <Input
              w="300px"
              size="lg"
              type="text"
              placeholder="Enter a title for your review"
              onChange={(e) => setTitle(e.target.value)}
              value={title}
            />
            <Box display="flex" alignItems="center" columnGap="2">
              <Text fontSize="18">Rating: </Text>
              <StarRatings
                starHoverColor="#F6E05E"
                starRatedColor="#F6E05E"
                starDimension="25px"
                rating={stars}
                numberOfStars={5}
                name="rating"
                isAggregateRating
                starSpacing="2px"
                changeRating={(rating) => {
                  setStars(rating);
                }}
              />
            </Box>
          </Box>
          <Textarea
            w="600px"
            h="120px"
            placeholder="Enter your review"
            onChange={(e) => setDesc(e.target.value)}
            value={desc}
          />
          <Box display="flex" justifyContent="flex-end">
            <Button w="200px" colorScheme="green" onClick={submitReview}>
              Submit
            </Button>
          </Box>
        </Box>
      </Box>
      {reviews.length === 0 ? (
        <Box my="6">
          <Heading>No reviews posted yet. Be the first!</Heading>
        </Box>
      ) : (
        <Box>
          {reviews.map((rev) => (
            <ReviewCard
              key={rev.id}
              aid={aid}
              rev={rev}
              setActivity={setActivity}
            />
          ))}
        </Box>
      )}
    </Box>
  );
};

const ReviewCard = ({ aid, rev, setActivity }) => {
  const [avatar, setAvatar] = useState("");
  const [username, setUsername] = useState("");
  const [editMode, setEditMode] = useState(false);
  const [editTitle, setEditTitle] = useState("");
  const [editRating, setEditRating] = useState(0);
  const [editDesc, setEditDesc] = useState("");
  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);

  const uid = window.sessionStorage.getItem("uid");

  useEffect(() => {
    fetchUserInfo(rev.user_id, setAvatar, setUsername);
  }, [rev.user_id]);

  const updateReview = async () => {
    setEditMode(false);
    await sendUpdateReview(
      rev.id,
      aid,
      uid,
      false,
      editTitle,
      editDesc,
      editRating,
      setNotifMsg,
      setIsError,
      setActivity
    );
  };

  return (
    <Box py="5">
      {notifMsg && (
        <Box
          px="6"
          py="2"
          borderRadius="13px"
          bgColor={isError ? "tomato" : "limegreen"}
        >
          <Text fontSize="xl">{notifMsg}</Text>
        </Box>
      )}
      <Box
        display="flex"
        justifyContent="flex-start"
        alignItems="center"
        columnGap="10"
        borderTop="1px solid white"
        my="4"
        py="4"
        borderBottom="1px solid white"
      >
        <Box pt="5">
          <Avatar src={`http://localhost:8080/avatars/${avatar}`} />
          <Text my="2">{username}</Text>
        </Box>
        <Box display="flex" justifyContent="space-between" w="60vw">
          <Box display="flex" flexDir="column" rowGap="2">
            {editMode ? (
              <Input
                type="text"
                size="lg"
                value={editTitle}
                onChange={(e) => setEditTitle(e.target.value)}
              />
            ) : (
              <Heading size="lg">{rev.title}</Heading>
            )}
            {editMode ? (
              <Box
                display="flex"
                justifyContent="flex-start"
                alignItems="center"
                columnGap="4"
              >
                <Text>Rating: </Text>
                <StarRatings
                  starHoverColor="#F6E05E"
                  starRatedColor="#F6E05E"
                  starDimension="18px"
                  rating={editRating}
                  numberOfStars={5}
                  name="rating"
                  isAggregateRating
                  changeRating={(rating) => setEditRating(rating)}
                  starSpacing="2px"
                />
              </Box>
            ) : (
              <StarRatings
                starHoverColor="#F6E05E"
                starRatedColor="#F6E05E"
                starDimension="18px"
                rating={rev.rating}
                numberOfStars={5}
                name="rating"
                isAggregateRating
                starSpacing="2px"
              />
            )}
            {editMode ? (
              <Textarea
                w="500px"
                h="200px"
                value={editDesc}
                onChange={(e) => setEditDesc(e.target.value)}
              />
            ) : (
              <Text mt="4" w="500px">
                {rev.description}
              </Text>
            )}
          </Box>

          <Box display="flex" justifyContent="center" alignItems="center">
            {editMode ? (
              <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
                columnGap="5"
              >
                <Button
                  colorScheme="green"
                  leftIcon={<CheckIcon />}
                  onClick={updateReview}
                >
                  Save
                </Button>
                <Button
                  colorScheme="blue"
                  leftIcon={<CloseIcon />}
                  onClick={() => {
                    setEditMode(false);
                    setEditTitle("");
                    setEditRating(0);
                    setEditDesc("");
                  }}
                >
                  Cancel
                </Button>
              </Box>
            ) : (
              <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
                columnGap="5"
              >
                <IconButton
                  colorScheme="teal"
                  aria-label="edit review"
                  icon={<EditIcon />}
                  onClick={() => {
                    setEditMode(true);
                    setEditTitle(rev.title);
                    setEditRating(rev.rating);
                    setEditDesc(rev.description);
                  }}
                />
                <IconButton
                  colorScheme="red"
                  aria-label="delete review"
                  icon={<DeleteIcon />}
                />
              </Box>
            )}
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

const MyMapComponent = withScriptjs(
  withGoogleMap((props) => {
    return (
      <GoogleMap defaultZoom={15} defaultCenter={props.latLng}>
        {props.isMarkerShown && <Marker position={props.latLng} />}
      </GoogleMap>
    );
  })
);

export default Activity;
