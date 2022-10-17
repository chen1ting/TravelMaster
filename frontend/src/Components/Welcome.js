import { validSessionGuard } from "../common/common";
import { sendGenerateItineraryReq } from "../api/api";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Box, Heading, Input, Text, Button } from "@chakra-ui/react";
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";
const Welcome = () => {
  const navigate = useNavigate();
  useEffect(() => {
    validSessionGuard(navigate, "/");
  }, []);
  const user = window.sessionStorage.getItem("username");
  const session_token = window.sessionStorage.getItem("session_token");
  const now = new Date();
  const [startDateTime, setStartDateTime] = useState(
    new Date(now.getFullYear(), now.getMonth(), now.getDay(), 15)
  );
  const [endDateTime, setEndDateTime] = useState(
    new Date(now.getFullYear(), now.getMonth(), now.getDay() + 3, 18)
  );
  const [cats, setCats] = useState(new Set());

  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);

  const generateItinerary = async () => {
    const content = await sendGenerateItineraryReq(
      session_token,
      startDateTime,
      endDateTime,
      cats
    );
    if (content == null) {
      setNotifMsg("Sorry! Something went wrong on our end...");
      setIsError(true);
      await new Promise((resolve) => setTimeout(resolve, 4000)); // 1 sec
      setIsError(false);
      setNotifMsg("");
      return;
    }

    setNotifMsg("Success! Your itinerary is generated.");
    setIsError(false);
    await new Promise((resolve) => setTimeout(resolve, 1250)); // 1 sec
    navigate(`/edit-itinerary/${content.id}`);
  };

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
  return (
    <Box
      display="flex"
      flexDir="column"
      alignItems="center"
      justifyContent="flex-start"
      h="94vh"
    >
      <Box mt="2%" textAlign="center">
        <Heading size="2xl" fontWeight="2" my="5">
          Hi {user}!
        </Heading>
        <Text fontSize="xl">
          Let's get started on planning your next visit to Singapore! 🇸🇬
        </Text>
      </Box>

      {notifMsg && (
        <Box
          mt="3"
          color={'white'}
          bgColor={isError ? "red.400" : "green.500"}
          py="2"
          px="6"
          borderRadius="15px"
        >
          <Text>{notifMsg}</Text>
        </Box>
      )}

      <Box
        mt={notifMsg ? "0" : "12"}
        mb="5"
        w="100%"
        display="flex"
        justifyContent="center"
        alignItems="center"
        py="5"
      >
        <Box display="flex" justifyContent="center" alignItems="center">
          <Text whiteSpace="nowrap" fontSize="3xl">
            You would be in Singapore from
          </Text>
          <DatePicker
            selected={startDateTime}
            showTimeSelect
            onChange={(date) => setStartDateTime(date)}
            customInput={
              <Button
                mt="1"
                mx="3"
                borderBottom="1px solid white"
                p="2"
                borderRadius="15px"
                bgColor="gray.200"
                whiteSpace="nowrap"
              >
                {startDateTime.toLocaleString()}
              </Button>
            }
          />

          <Text whiteSpace="nowrap" fontSize="3xl">
            to
          </Text>
          <DatePicker
            selected={endDateTime}
            showTimeSelect
            onChange={(date) => setEndDateTime(date)}
            customInput={
              <Button
                mt="1"
                ml="3"
                mr="1"
                borderBottom="1px solid white"
                p="2"
                borderRadius="15px"
                bgColor="gray.200"
                whiteSpace="nowrap"
              >
                {endDateTime.toLocaleString()}
              </Button>
            }
          />
          <Text whiteSpace="nowrap" fontSize="3xl">
            .
          </Text>
        </Box>
      </Box>
      <Box textAlign="center">
        <Text fontSize="18">
          Select a few categories you would be interested in:
        </Text>
        <Box
          display="flex"
          flexWrap="wrap"
          padding="5"
          maxW="500px"
          justifyContent="center"
          alignItems="center"
          columnGap="7"
          rowGap="6"
        >
          {catsList.map((cat) => (
            <Button
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
        </Box>
      </Box>

      <Box my="10" w="40%" textAlign="center">
        <Button
          colorScheme="whatsapp"
          w="full"
          h="12"
          onClick={generateItinerary}
        >
          <Text fontSize="18">Plan my itinerary!</Text>
        </Button>
      </Box>
      <Box p={'1vh'}></Box>
    </Box>
  );
};

export default Welcome;
