import {
    Avatar,
    Box,
    Grid,
    GridItem,
    Tab,
    TabList,
    TabPanel,
    TabPanels,
    Tabs,
    Text,
} from "@chakra-ui/react";
import {ENDPOINT} from "../api/api";
import ActivityFeed from "../Components/ProfileSubPages/ActivityFeed";
import Itineraries from "../Components/ProfileSubPages/Itineraries";
import {useEffect, useState} from "react";
import {getUserActivities} from "../api/api";
import ReviewsPosted from "../Components/ProfileSubPages/ReviewsPosted";

const Profile = ({imageUrl}) => {
    const user = window.sessionStorage.getItem("username");
    const uid = window.sessionStorage.getItem("uid");
    const [isLoading, setIsLoading] = useState(false);
    const [activities, setActivities] = useState(false);
    const [reviews, setReviews] = useState(false);

    useEffect(() => {
        getUserActivities(uid, setIsLoading, setActivities, setReviews);
    }, []);

    return (
        <Grid
            templateAreas={`'spacing spacing'
                            'icon username'
                            'tabs tabs'`}
            gridTemplateRows={"2% 15% 83%"}
            gridTemplateColumns={"18% 82%"}
            h="100vh"
            ml={"2%"}
            mr={"2%"}
        >
            <GridItem area={'spacing'}>
                <Box p={'1px'}></Box>
            </GridItem>
            <GridItem area={"icon"}>
                <Box
                    position={"relative"}
                    top={"50%"}
                    left={"50%"}
                    transform={"translate(-50%,-50%)"}
                    textAlign={"center"}
                >
                    <Avatar size={"2xl"} src={`${ENDPOINT}/avatars/` + imageUrl}/>
                </Box>
            </GridItem>
            <GridItem area={"username"}>
                <Box
                    position={"relative"}
                    top={"50%"}
                    left={"50%"}
                    transform={"translate(-50%,-50%)"}
                >
                    {/*Add username here*/}
                    <Text fontSize="4xl">{user}</Text>
                </Box>
            </GridItem>
            {/* <GridItem area={"about"} bg={"lightskyblue"}></GridItem> */}
            <GridItem area={"tabs"}>
                <Tabs isFitted variant="enclosed" mt={'1%'} ml={'5%'} mr={'5%'}>
                    <TabList>
                        <Tab
                            borderColor={'gray.400'}
                            _selected={{color: "white", bg: "orange.500"}}
                            fontSize={'xl'}
                        >
                            Activities Created
                        </Tab>
                        <Tab
                            borderColor={'gray.400'}
                            _selected={{color: "white", bg: "green.500"}}
                            fontSize={'xl'}
                        >
                            Reviews Posted
                        </Tab>
                        <Tab
                            borderColor={'gray.400'}
                            _selected={{color: "white", bg: "blue.500"}}
                            fontSize={'xl'}
                        >
                            Itineraries
                        </Tab>
                    </TabList>
                    <TabPanels>
                        <TabPanel>
                            <ActivityFeed></ActivityFeed>
                        </TabPanel>
                        <TabPanel>
                            <ReviewsPosted></ReviewsPosted>
                        </TabPanel>
                        <TabPanel>
                            <Itineraries></Itineraries>
                        </TabPanel>
                    </TabPanels>
                </Tabs>
            </GridItem>
        </Grid>
    );
};

export default Profile;
