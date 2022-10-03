import {Box, Grid, GridItem, Stack, HStack, Text} from "@chakra-ui/react";

const BookingItems = () => {
    const BookingItem = {
        img: 'https://bit.ly/2Z4KKcF',
        title: 'Activity Title',
        address:'Activity Address',
        desc: 'Activity Description',
        booked: true
    }

    return (
        <HStack>
            {BookingItem.img}
            <Stack maxW='sm' borderWidth='1px' borderRadius='lg' overflow='hidden' justify='center' align='center'>
                <Box>
                    {BookingItem.title}
                </Box>
                <Box>
                    {BookingItem.address}
                </Box>
                <Box>
                    {BookingItem.desc}
                </Box>
            </Stack>
        </HStack>

    )
}

const Bookings = () => {
    return (
    <Grid
        templateRows='repeat(3, 1fr)'
        gap={4}
    >
        <GridItem rowSpan={1} bg='azure' justify="center" >
            <BookingItems />
        </GridItem>
        <GridItem rowSpan={1} bg='lightcyan' justify = 'center' alignContent='center'>
            <BookingItems />
        </GridItem>
        <GridItem rowSpan={1} bg='azure'>
            <BookingItems />
        </GridItem>
    </Grid>
)
}

export default Bookings;