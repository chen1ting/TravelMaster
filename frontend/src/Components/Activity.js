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
} from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { getActivityById } from "../api/api";
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
        <Box
          display="flex"
          justifyContent="center"
          mt="12"
          py="10"
          columnGap="10%"
          maxW="80%"
        >
          <Box>
            <Image
              w="650px"
              h="300px"
              src={`http://localhost:8080/activity-images/${act.image_names[0]}`}
              alt={act.title}
              borderRadius="20px"
            />
            <Box mt="8" mb="2" display="flex" alignItems="center" columnGap="2">
              {act.category.map((cat) => (
                <Badge variant="outline" colorScheme="green">
                  {cat}
                </Badge>
              ))}
            </Box>
            <Badge my="4" size="lg" colorScheme={act.paid ? "yellow" : "teal"}>
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
              <Table size="sm">
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
      )}
    </Box>
  );
};

export default Activity;
