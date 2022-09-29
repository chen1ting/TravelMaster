import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { validSessionGuard } from "../common/common";
import { getItinerary } from "../api/api";
import { useNavigate } from "react-router-dom";
import { Box, Heading, Text } from "@chakra-ui/react";

const EditItinerary = () => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(true);
  const [notifMsg, setNotifMsg] = useState("");
  const [timeBins, setTimeBins] = useState([]);
  const [itineraryMap, setItineraryMap] = useState(new Map());
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const session_token = window.sessionStorage.getItem("session_token");

  const { id } = useParams();
  useEffect(() => {
    validSessionGuard(navigate);
    getItinerary(
      parseInt(id),
      session_token,
      setIsLoading,
      setNotifMsg,
      setTimeBins,
      setItineraryMap,
      setStartDate,
      setEndDate
    );
  }, [id, navigate, session_token]);

  return (
    <Box>
      {isLoading ? (
        <Text>Loading ... </Text>
      ) : notifMsg ? (
        <Text>{notifMsg}</Text>
      ) : (
        <Box>
          <Heading>Hello world</Heading>
          {timeBins[31].name} {timeBins[31].description}
        </Box>
      )}
    </Box>
  );
};

export default EditItinerary;
