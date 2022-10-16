import {
  Flex,
  Box,
  Grid,
  Stack,
  Button,
  Heading,
  Text,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ActivityCard } from "../../Pages/Discover";
import { ReviewCard } from "../Activity";
import { fetchProfile } from "../../api/api";

const ActivityFeed = () => {
  const [reviews, setReviews] = useState([]);
  const [activities, setActivities] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    fetchProfile(
      parseInt(window.sessionStorage.getItem("uid")),
      setReviews,
      setActivities,
      setIsLoading
    );
  }, []);

  const navigate = useNavigate();

  return (
    <Box display="flex" flexDir="column">
      {isLoading ? (
        <Text>Loading...</Text>
      ) : (
        <Box display="flex" flexDir="column">
          <Heading mt="12">Reviews</Heading>
          {reviews &&
            reviews.map((review) => (
              <ReviewCard
                aid={review.aid}
                rev={review}
                setActivities={() => {}}
              />
            ))}
          <Heading mt="12">Activities created</Heading>
          {activities &&
            activities.map((act) => (
              <ActivityCard act={act} navigate={navigate} />
            ))}
        </Box>
      )}
    </Box>
  );
};

export default ActivityFeed;
