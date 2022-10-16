import { Box, Image, Link } from "@chakra-ui/react";
import { StarIcon } from "@chakra-ui/icons";
import { useNavigate } from "react-router-dom";

const ActivityCard = (
  activityID,
  title,
  rating,
  paidActivity,
  activityCategories,
  imageUrl
) => {
  const navigate = useNavigate();

  return (
    <Box
      maxW="sm"
      borderWidth="1px"
      borderRadius="lg"
      overflow="hidden"
      bgColor="orange.50"
    >
      <Image src={imageUrl} />

      <Box p="6">
        <Box fontSize="lg" fontWeight="bold" lineHeight="tight" noOfLines={2}>
          <Link
            onClick={() => {
              navigate("/activitydescription"); // Change this
            }}
          >
            {title}
          </Link>
        </Box>

        <Box as="span" color="gray.600" noOfLines={1}>
          {paidActivity ? "Paid Activity" : "FREE!"}
        </Box>

        <Box as="span" color="gray.600" noOfLines={1}>
          {activityCategories}
        </Box>

        <Box display="flex" mt="2" alignItems="center">
          {Array(5)
            .fill("")
            .map((_, i) => (
              <StarIcon key={i} color={i < rating ? "teal.500" : "gray.300"} />
            ))}
          <Box as="span" ml="2" color="gray.600" fontSize="sm" href="reviews">
            <Link
              onClick={() => {
                navigate("/reviews"); // Change this
              }}
            >
              Activity Reviews
            </Link>
          </Box>
        </Box>
      </Box>
    </Box>
  );
};

export default ActivityCard;
