import { useState, useEffect } from "react";
import { validSessionGuard } from "../../common/common";
import { getItisByUser } from "../../api/api";
import {
  Button,
  Center,
  Flex,
  Grid,
  GridItem,
  Spacer,
  Text,
  Box,
  Heading,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import DisplayItineraries from "../ItinerariesComponents/DisplayItineraries";
import { PlusSquareIcon } from "@chakra-ui/icons";
import ItineraryCard from "../ItinerariesComponents/ItineraryCard";

const Itineraries = () => {
  const navigate = useNavigate();
  const [itis, setItis] = useState([]);
  useEffect(() => {
    validSessionGuard(navigate, "/");
    getItisByUser(window.sessionStorage.getItem("session_token"), setItis);
  }, [navigate]);

  return (
    <Box
      display="flex"
      flexDir="column"
      alignItems="center"
      w="100%"
      rowGap="5"
    >
      <Box
        mt="10"
        display="flex"
        justifyContent="center"
        alignItems="center"
        w="full"
      >
        <Heading>My Itineraries</Heading>
      </Box>
      <Box display="flex" justifyContent="flex-end" w="full" mr="20%">
        <Button
          leftIcon={<PlusSquareIcon />}
          onClick={() => {
            navigate("/welcome");
          }}
          w={250}
          colorScheme="teal"
        >
          <Text size="sm">Create New Itinerary</Text>
        </Button>
      </Box>
      <Box
        display="flex"
        justifyContent="space-evenly"
        alignItems="center"
        flexWrap="wrap"
        rowGap="10"
        columnGap="5"
      >
        {itis.map((it) => (
          <ItineraryCard it={it} />
        ))}
      </Box>
    </Box>
  );
};

export default Itineraries;
