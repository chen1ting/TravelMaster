import {
    Box,
    Grid,
    GridItem,
    Tab, TabList, TabPanel, TabPanels, Tabs,
    Text
} from "@chakra-ui/react";
import DiscoverActivities from "./DiscoverSubPages/DiscoverActivities";
import DiscoverFlights from "./DiscoverSubPages/DiscoverFlights";
import DiscoverHotels from "./DiscoverSubPages/DiscoverHotels";

const DiscoverDeprecated = () => {
    return (
        <Grid
            templateAreas={`"title"
                            "content"
                            `}
            gridTemplateRows={'10fr 90fr'}
            h='100vh'
            fontWeight='bold'
            bgColor={'blue.50'}
        >
            <GridItem area={"title"}>
                <Box position={'relative'} top={'23%'} left={'7.5%'} transform={'translate(-50%,-50%)'}
                     textAlign={"left"} width='fit-content' mt={'1%'}>
                    <Text fontSize='4xl'>Discover</Text>
                </Box>
            </GridItem>
            <GridItem area={"content"} bgColor={'white'}>
                <Tabs isFitted variant='enclosed'>
                    <TabList>
                        <Tab _selected={{color: 'white', bg: 'orange.500'}} fontWeight={'bold'} fontSize='lg'>Events & Activities</Tab>
                        <Tab _selected={{color: 'white', bg: 'green.400'}} fontWeight={'bold'} fontSize='lg'>Flights</Tab>
                        <Tab _selected={{color: 'white', bg: 'blue.400'}} fontWeight={'bold'} fontSize='lg'>Hotels</Tab>
                    </TabList>
                    <TabPanels>
                        <TabPanel>
                            <DiscoverActivities></DiscoverActivities>
                        </TabPanel>
                        <TabPanel>
                            <DiscoverFlights></DiscoverFlights>
                        </TabPanel>
                        <TabPanel>
                            <DiscoverHotels></DiscoverHotels>
                        </TabPanel>
                    </TabPanels>
                </Tabs>
            </GridItem>
        </Grid>
    );
}

export default DiscoverDeprecated;