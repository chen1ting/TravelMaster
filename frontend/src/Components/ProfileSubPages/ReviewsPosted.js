import {
    Box, Center,
    Text,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { ReviewCard } from "../Activity";
import { fetchProfile } from "../../api/api";

const ReviewsPosted = () => {
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

    return (
        <Center display="flex" flexDir="column">
            {isLoading ? (
                <Text>Loading...</Text>
            ) : (
                <Box display="flex" flexDir="column">
                    {reviews &&
                        reviews.map((review) => (
                            <ReviewCard
                                aid={review.aid}
                                rev={review}
                                setActivities={() => {}}
                            />
                        ))}
                </Box>
            )}
        </Center>
    );
};

export default ReviewsPosted;
