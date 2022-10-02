import {Box, Image, Link} from "@chakra-ui/react";
import {StarIcon} from "@chakra-ui/icons";
import {useNavigate} from "react-router-dom";

const ActivityCard = (activityID, title, rating, paidActivity, activityCategory, imageUrl) => {
    const navigate = useNavigate();

    return (<Box maxW='sm' borderWidth='1px' borderRadius='lg' overflow='hidden' bgColor='orange.50'>
        <Image src={imageUrl}/>

        <Box p='6'>
            <Box
                fontSize='lg'
                fontWeight='bold'
                lineHeight='tight'
                noOfLines={2}
            >
                <Link href="frontend/src/Pages/ActivityDescription">{title}</Link>
            </Box>

            <Box as='span' color='gray.600'>
                {paidActivity ? 'Paid Activity' : 'FREE!'}
            </Box>

            <Box as='span' color='gray.600'>
                {activityCategory}
            </Box>

            <Box display='flex' mt='2' alignItems='center'>
                {Array(5)
                    .fill('')
                    .map((_, i) => (<StarIcon
                        key={i}
                        color={i < rating ? 'teal.500' : 'gray.300'}
                    />))}
                <Box as='span' ml='2' color='gray.600' fontSize='sm' href="reviews">
                    <Link
                        onClick={ () => {
                            navigate("frontend/src/Pages/Reviews");
                        }}
                    >Activity Reviews</Link>
                </Box>
            </Box>
        </Box>
    </Box>)
}

export default ActivityCard;