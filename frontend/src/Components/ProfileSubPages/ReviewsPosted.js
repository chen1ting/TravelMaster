import { Box, Center, Text } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { ReviewCard } from "../Activity";
import { fetchProfile } from "../../api/api";

const ReviewsPosted = ({ isLoading, reviews, setReviews }) => {
  return (
    <Center display="flex" flexDir="column">
      {isLoading ? (
        <Text>Loading...</Text>
      ) : (
        <Box display="flex" flexDir="column">
          {reviews &&
            reviews.map((review) => (
              <ReviewCard
                aid={review.activity_id}
                rev={review}
                setActivity={(act) => {
                  setReviews((r) =>
                    r.map((prev) =>
                      act.review_list.find((rev) => rev.id === review.id)
                        ? act.review_list.find((rev) => rev.id === review.id)
                        : prev
                    )
                  );
                }}
              />
            ))}
        </Box>
      )}
    </Center>
  );
};

export default ReviewsPosted;
