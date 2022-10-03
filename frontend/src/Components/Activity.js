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
} from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { getActivityById, addReview } from "../api/api";
import StarRatings from "react-star-ratings";

const Activity = () => {
  const { id } = useParams();
  const [isLoading, setIsLoading] = useState(true);
  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);
  const [act, setActivity] = useState(null);

  useEffect(() => {
    getActivityById(id, setActivity, setIsLoading);
  }, []);

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
            columnGap="10%"
          >
            <Box>
              <Image
                w="600px"
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
                <Text mt="1px">({act.rating_score})</Text>
              </Box>
            </Box>
            <Box>
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
    } else if (code != 201) {
      setIsError(true);
      setNotifMsg("Something went wrong...");
      await new Promise((resolve) => setTimeout(resolve, 3000));
      setIsError(false);
      setNotifMsg("");
      return;
    }

    setIsError(false);
    setNotifMsg("Successfully uploaded your review.");
    await new Promise((resolve) => setTimeout(resolve, 3000));
    setNotifMsg("");
  };
  return (
    <Box>
      <Heading mt="3" mb="5">
        Reviews
      </Heading>
      {notifMsg && (
        <Text
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
            w="700px"
            h="200px"
            placeholder="Enter your review"
            onChange={(e) => setDesc(e.target.value)}
          />
          <Box display="flex" justifyContent="flex-end">
            <Button w="200px" colorScheme="green" onClick={submitReview}>
              Submit
            </Button>
          </Box>
        </Box>
      </Box>
      {reviews.length === 0 ? (
        <Box>
          <Heading>No reviews posted yet. Be the first!</Heading>
        </Box>
      ) : (
        <Box>
          {reviews.map((rev) => (
            <Box>
              <Text>{rev.title}</Text>
            </Box>
          ))}
        </Box>
      )}
    </Box>
  );
};

export default Activity;