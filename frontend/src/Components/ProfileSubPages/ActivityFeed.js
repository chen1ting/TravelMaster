import { Box, Center, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ActivityCard } from "../../Pages/Discover";
import { fetchProfile } from "../../api/api";

const ActivityFeed = ({ isLoading, activities }) => {
  const navigate = useNavigate();

  return (
    <Box display="flex" flexDir="column">
      {isLoading ? (
        <Text>Loading...</Text>
      ) : (
        <Center display="flex" flexDir="column">
          {activities &&
            activities.map((act) => (
              <ActivityCard
                act={{ id: act.activity_id, ...act }}
                navigate={navigate}
              />
            ))}
        </Center>
      )}
    </Box>
  );
};

export default ActivityFeed;
