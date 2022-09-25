import {Box, Grid, GridItem, Tab, TabList, TabPanel, TabPanels, Tabs, Text} from "@chakra-ui/react";
import ActivityFeed from "../Components/ProfileSubPages/ActivityFeed";
import Itineraries from "../Components/ProfileSubPages/Itineraries";
import Bookings from "../Components/ProfileSubPages/Bookings";

const Profile = () => {
    return (
        <Grid
            templateAreas={`'icon username username'
                            'about about tabs'`}
            gridTemplateRows={'15% 85%'}
            gridTemplateColumns={'10% 10% 80%'}
            h='100vh'
            mt={'1%'}
            ml={'3%'}
            mr={'3%'}
        >
            <GridItem area={'icon'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    {/*Add user icon here*/}
                    Icon
                </Box>

            </GridItem>
            <GridItem area={'username'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}>
                    {/*Add username here*/}
                    <Text fontSize='4xl'>Username123</Text>
                </Box>
            </GridItem>
            <GridItem area={'about'} bg={'lightskyblue'}>
                <Box
                    mt={'10%'}
                    ml={'10%'}
                    mr={'10%'}>
                    <Text fontWeight='bold'>About you</Text>
                    !!! Add about you here !!!
                </Box>
            </GridItem>
            <GridItem area={'tabs'}>
                <Tabs isFitted variant='enclosed'>
                    <TabList mb='1em'>
                        <Tab _selected={{ color: 'white', bg: 'orange.500' }}>Activity Feed</Tab>
                        <Tab _selected={{ color: 'white', bg: 'green.400' }}>Itineraries</Tab>
                        <Tab _selected={{ color: 'white', bg: 'blue.400' }}>Bookings</Tab>
                    </TabList>
                    <TabPanels>
                        <TabPanel>
                            <ActivityFeed></ActivityFeed>
                        </TabPanel>
                        <TabPanel>
                            <Itineraries></Itineraries>
                        </TabPanel>
                        <TabPanel>
                            <Bookings></Bookings>
                        </TabPanel>
                    </TabPanels>
                </Tabs>
            </GridItem>
        </Grid>
    )
}

export default Profile;