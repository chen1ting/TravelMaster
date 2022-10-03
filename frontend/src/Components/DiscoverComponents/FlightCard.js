import {Box, Link} from "@chakra-ui/react";

const FlightCard = () => {
    let currentDate = new Date();
    const property = {
        airlines: 'Singapore Airlines',
        departureTime: currentDate.getHours() + ':' + currentDate.getMinutes(),
        arrivalTime: currentDate.getHours() + ':' + currentDate.getMinutes(),
        departureLocation: 'Tokyo',
        arrivalLocation: 'Singapore Changi Airport',
        ticketPrice: '$1000',
        numberStops: 0,
        url: 'https://www.expedia.com.sg/',
    }

    return (
        <Box maxW='md' borderWidth='1px' borderRadius='lg' overflow='hidden' bgColor='teal.100'>
            <Box p='6'>
                <Box
                    fontWeight='bold'
                    letterSpacing='wide'
                    fontSize='lg'
                    textTransform='uppercase'
                >
                    {property.airlines}
                </Box>
                <Box
                    fontWeight='bold'
                    letterSpacing='wide'
                    fontSize='lg'
                    textTransform='uppercase'
                >
                    {property.departureTime} - {property.arrivalTime}
                </Box>
                <Box
                    mt='1'
                    fontWeight='semibold'
                    as='h4'
                    lineHeight='tight'
                    noOfLines={1}
                >
                    <Link href="frontend/src/Pages/ActivityDescription">{property.title}</Link>
                </Box>
                <Box
                    fontSize='md'
                    textTransform='uppercase'
                >
                    {property.departureLocation}
                </Box>
                <Box
                    fontSize='md'
                >
                    to
                </Box>
                <Box
                    fontSize='md'
                    textTransform='uppercase'
                >
                    {property.arrivalLocation}
                </Box>
                <Box
                    fontSize='2xl'>
                    {property.ticketPrice}
                    <Box as='span' color='gray.600' fontSize='sm'>
                        {` per traveller`}
                    </Box>
                </Box>
            </Box>
        </Box>
    )
}

export default FlightCard;