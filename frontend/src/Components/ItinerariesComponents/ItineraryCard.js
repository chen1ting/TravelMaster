import { useNavigate } from "react-router-dom";
import { Box, Heading, Link, Text } from "@chakra-ui/react";

const ItineraryCard = ({ it }) => {
  const navigate = useNavigate();

  return (
    <Box
      w="500px"
      border="1px solid white"
      p="5"
      borderRadius="20px"
      textAlign="center"
      cursor="pointer"
      onClick={() => navigate(`/edit-itinerary/${it.id}`)}
      _hover={{ bg: "teal.400" }}
    >
      <Heading size="md">{it.name}</Heading>
      <Box mt="3" display="flex" justifyContent="space-evenly">
        <Text>From: {new Date(it.start_time * 1000).toLocaleDateString()}</Text>
        <Text>To: {new Date(it.end_time * 1000).toLocaleDateString()}</Text>
      </Box>
    </Box>
  );
};

export default ItineraryCard;
