import {useNavigate} from "react-router-dom";
import {Box, Link} from "@chakra-ui/react";

const ItineraryCard = (itineraryID, title, startDate, endDate) => {
    const navigate = useNavigate();

    return (
        <Box maxW='sm' borderWidth='1px' borderRadius='lg' overflow='hidden' bgColor='orange.50'>
            <Box p='6'>
                <Box
                    fontSize='lg'
                    fontWeight='bold'
                    lineHeight='tight'
                    noOfLines={2}
                >
                    <Link
                        onClick={() => {
                            navigate("/"); // Change this
                        }}
                    >{title}</Link>
                </Box>

                <Box as='span' color='gray.600' noOfLines={1}>
                    {startDate} {`-`} {endDate}
                </Box>

            </Box>
        </Box>
    )
}

export default ItineraryCard;