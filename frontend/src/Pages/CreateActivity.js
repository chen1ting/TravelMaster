import {useState} from 'react';
//import { useAuth } from '../lib/auth';
import Geocode from "react-geocode";

import {
    FormControl, FormLabel, Button, Flex, Grid, GridItem,
    AspectRatio,
    Box,
    BoxProps,
    Container,
    forwardRef,
    Heading,
    Input,
    HStack,
    VStack,
    StackDivider,
    Stack,
    Text,
    Checkbox, CheckboxGroup, Spacer
} from "@chakra-ui/react";

import { motion, useAnimation } from "framer-motion";
import { sendCreateActivityReq } from "../api/apiCreateActivity";
import {useNavigate} from "react-router-dom";
const fields_width = '52.5%';


// set response language. Defaults to english.
Geocode.setLanguage("en");

// set response region. Its optional.
// A Geocoding request with region=es (Spain) will return the Spanish city.
Geocode.setRegion("sg");

// set location_type filter . Its optional.
// google geocoder returns more that one address for given lat/lng.
// In some case we need one address as response for which google itself provides a location_type filter.
// So we can easily parse the result for fetching address components
// ROOFTOP, RANGE_INTERPOLATED, GEOMETRIC_CENTER, APPROXIMATE are the accepted values.
// And according to the below google docs in description, ROOFTOP param returns the most accurate result.
Geocode.setLocationType("ROOFTOP");

// Enable or disable logs. Its optional.
Geocode.enableDebug();


const CreateActivity = () => {
    const [descriptionlocation, setDescriptionLocation] = useState('');
    const [addressactivity, setAddressActivity] = useState('');
    const [descriptionActivity, setDescriptionActivity] = useState('');
    const [activityname, setActivityName] = useState('');
    const [ispaid, setIsPaid] = useState('');
    const [sundayopenhr, setSundayOpenHr] = useState('');
    const [sundayclosehr, setSundayCloseHr] = useState('');
    const [mondayopenhr, setMondayOpenHr] = useState('');
    const [mondayclosehr, setMondayCloseHr] = useState('');
    const [tuesdayopenhr, setTuesdayOpenHr] = useState('');
    const [tuesdayclosehr, setTuesdayCloseHr] = useState('');
    const [wednesdayopenhr, setWednesdayOpenHr] = useState('');
    const [wednesdayclosehr, setWednesdayCloseHr] = useState('');
    const [thursdayopenhr, setThursdayOpenHr] = useState('');
    const [thursdayclosehr, setThursdayCloseHr] = useState('');
    const [fridayopenhr, setFridayOpenHr] = useState('');
    const [fridayclosehr, setFridayCloseHr] = useState('');
    const [saturdayopenhr, setSaturdayOpenHr] = useState('');
    const [saturdayclosehr, setSaturdayCloseHr] = useState('');

    const [image, setImage] = useState('');

    const [showError, setShowError] = useState(false);
    const [errMsg, setErrMsg] = useState("");
    const navigate = useNavigate();

    async function onSubmit(e) {
        e.preventDefault();
        // might wanna consider adding a regex check for email format
        // and also password validation regex
        var bad =
            descriptionlocation === "" ||
            addressactivity === "" ||
            descriptionActivity === "" ||
            activityname === "" ||
        setShowError(bad);
        if (bad) {
            setErrMsg("A valid review and title is required.");
            return;
        }
        setErrMsg(""); // always clear after

        const data = await sendCreateActivityReq(descriptionlocation, addressactivity, descriptionActivity, activityname, ispaid, image, sundayopenhr
            , sundayclosehr, mondayopenhr, mondayclosehr, tuesdayopenhr, tuesdayclosehr, wednesdayopenhr, wednesdayclosehr
            , thursdayopenhr, thursdayclosehr, fridayopenhr, fridayclosehr, saturdayopenhr, saturdayclosehr); //////TO CHANGE THE FUNCTION
        if (data == null) {
            setShowError(true);
            setErrMsg("Sorry, something went wrong on our side.");
            return;
        }


        // redirect to homepage
        navigate("/");
    }

    return (
        <Grid
            templateAreas={`"left_top right"
                            "left_bottom right"
                            `}
            gridTemplateRows={'20fr 80fr'}
            gridTemplateColumns={'1fr 2fr'}
            h='100vh'
            // gap='1'
            // color='blackAlpha.700'
            fontWeight='bold'
        >
            <GridItem pl='2' bg='blue.50' area={'left_top'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    <Text fontSize='2xl'>Activity</Text>
                </Box>

            </GridItem>

            <GridItem pl='2' bg='blue.100' area={'left_bottom'}>
                <VStack
                    spacing={4}
                    align='center'>
                    <Spacer/>
                    <Input size="lg" type="file" onChange={(e) => setImage(e.target.value)}/>
                    <Input
                        m={4}
                        w={fields_width}
                        bgColor={'whitesmoke'}
                        type="text"
                        placeholder="Name of Event"
                        onChange={(e) => setActivityName(e.target.value)}
                    ></Input>
                    <Checkbox onChange={(e) => setIsPaid(e.target.value)}>Is it an ticketed entrance?</Checkbox>
                </VStack>
            </GridItem>



            <GridItem pl='2' bg='white' area={'right'}>
                <VStack
                    divider={<StackDivider borderColor='gray.200' />}
                    spacing={10}
                    align="normal"
                    >       
                            <Spacer/>
                            <Flex
                                wrap="wrap"
                                direction="column"
                                align="start"
                            >
                                <Stack spacing={8}>
                                    <HStack spacing='24px'>
                                        <Text fontSize='xl'>Description of Location</Text>
                                        <Input
                                            m={4}
                                            w={"96"}
                                            bgColor={'whitesmoke'}
                                            type="text"
                                            placeholder=""
                                            onChange={(e) => setDescriptionLocation(e.target.value)}
                                        ></Input>
                                    </HStack>
                                    <HStack spacing='86px'>
                                        <Text fontSize='xl'>Address of Activity</Text>
                                        <Input
                                            m={4}
                                            w={"96"}
                                            bgColor={'whitesmoke'}
                                            type="text"
                                            placeholder=""
                                            onChange={(e) => setAddressActivity(e.target.value)}
                                        ></Input>
                                    </HStack>
                                    <HStack spacing='54px'>
                                        <Text fontSize='xl'>Description of Activity</Text>
                                        <Input
                                            m={4}
                                            w={"96"}
                                            h={"56"}
                                            bgColor={'whitesmoke'}
                                            type="text"
                                            placeholder=""
                                            onChange={(e) => setDescriptionActivity(e.target.value)}
                                        ></Input>
                                    </HStack>
                                </Stack>
                                <Spacer/>
                            </Flex>
                            <Flex
                                wrap="wrap"
                                direction="column"
                                align="left"
                            >

                                <Stack spacing={4}>
                                    <Text fontSize='3xl'>Activity Hours</Text>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl' size="2000">Sunday</Text>
                                        <Input size="md" type="time" onChange={(e) => setSundayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setSundayCloseHr(e.target.value)}/>
                                    </HStack>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl'>Monday</Text>
                                        <Input size="md" type="time" onChange={(e) => setMondayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setMondayCloseHr(e.target.value)}/>
                                    </HStack>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl'>Tuesday</Text>
                                        <Input size="md" type="time" onChange={(e) => setTuesdayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setTuesdayCloseHr(e.target.value)}/>
                                    </HStack>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl'>Wednesday</Text>
                                        <Input size="md" type="time" onChange={(e) => setWednesdayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setWednesdayCloseHr(e.target.value)}/>
                                    </HStack>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl'>Thursday</Text>
                                        <Input size="md" type="time" onChange={(e) => setThursdayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setThursdayCloseHr(e.target.value)}/>
                                    </HStack>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl'>Friday</Text>
                                        <Input size="md" type="time" onChange={(e) => setFridayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setFridayCloseHr(e.target.value)}/>
                                    </HStack>
                                    <HStack spacing='24px'>
                                        <Text fontSize='1xl'>Saturday</Text>
                                        <Input size="md" type="time" onChange={(e) => setSaturdayOpenHr(e.target.value)}/>
                                        <Text fontSize='1xl'> to </Text>
                                        <Input size="md" type="time" onChange={(e) => setSaturdayCloseHr(e.target.value)}/>
                                    </HStack>
                                </Stack>
                            </Flex>
                    <Button w={fields_width} m={4} type="submit" onClick={onSubmit} bg="teal">
                        <font size={5} color={'white'}>Create Event</font>
                    </Button>
                            <Spacer/>

                </VStack>
            </GridItem>
        </Grid>

    );
};

export default CreateActivity;