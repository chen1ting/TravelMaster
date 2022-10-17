import {useState} from 'react';
//import { useAuth } from '../lib/auth';
import {
    Button, Flex, Grid, GridItem,
    AspectRatio,
    Box,
    Heading,
    Input,
    HStack,
    VStack,
    StackDivider,
    Stack,
    Text,
    Checkbox, Spacer
} from "@chakra-ui/react";

import { motion } from "framer-motion";
import {sendCreateItineraryReq} from "../api/apiCreateItinerary";
import { useNavigate } from "react-router-dom";

//import UploadImage from './src/Components/UploadImage';
const fields_width = '52.5%';

const CreateItinerary = () => {
    const [descriptionlocation, setDescriptionLocation] = useState('');
    const [addressactivity, setAddressActivity] = useState('');
    const [descriptionactivity, setDescriptionActivity] = useState('');
    const [activityname, setActivityname] = useState('');
    const [visitdatestart, setVisitDateStart] = useState('');
    const [visitdateend, setVisitDateEnd] = useState('');
    const navigate = useNavigate();
    const [showError, setShowError] = useState(false);
    const [errMsg, setErrMsg] = useState("");

    //const { signIn } = useAuth();

    async function onSubmit(e) {
        e.preventDefault();
        // might wanna consider adding a regex check for email format
        // and also password validation regex
        var bad =
            descriptionlocation === "" ||
            addressactivity === "" ||
            descriptionactivity === "" ||
            activityname === "" ||
            visitdatestart === "" ||
            visitdateend === "";
        setShowError(bad);
        if (bad) {
            if (descriptionlocation === "") {
                setErrMsg("A valid description location is required.");
            } else if (addressactivity === "") {
                setErrMsg("A valid activity address is required.");
            } else if (descriptionactivity === "") {
                setErrMsg("A valid activity description is required.");
            } else if (activityname === "") {
                setErrMsg("A valid activity name is required.");
            } else if (visitdatestart === "") {
                setErrMsg("A valid start date is required.");
            } else {
                setErrMsg("A valid end date is required.");
            }
            return;
        }
        setErrMsg(""); // always clear after

        const data = await sendCreateItineraryReq(descriptionlocation, addressactivity, descriptionactivity, activityname, visitdatestart, visitdateend);
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
                    <Text fontSize='2xl'>Create Itinerary</Text>
                </Box>

            </GridItem>



            <GridItem pl='2' bg='blue.100' area={'left_bottom'}>
                <VStack
                    spacing={4}
                    align='center'>
                    <Spacer/>
                    <AspectRatio width="64" ratio={1}>
                        <Box
                            borderColor="gray.300"
                            borderStyle="dashed"
                            borderWidth="2px"
                            rounded="md"
                            shadow="sm"
                            role="group"
                            transition="all 150ms ease-in-out"
                            _hover={{
                                shadow: "md"
                            }}
                            as={motion.div}
                            initial="rest"
                            animate="rest"
                            whileHover="hover"
                            textAlign={"center"}
                        >
                            <Box position="relative" height="100%" width="100%">
                                <Box
                                    position="absolute"
                                    top="0"
                                    left="0"
                                    height="100%"
                                    width="100%"
                                    display="flex"
                                    flexDirection="column"
                                >
                                    <Stack
                                        height="100%"
                                        width="100%"
                                        display="flex"
                                        alignItems="center"
                                        justify="center"
                                        spacing="4"
                                    >
                                        <Box height="16" width="12" position="relative">

                                        </Box>
                                        <Stack p="8" textAlign="center" spacing="1">
                                            <Heading fontSize="lg" color="gray.700" fontWeight="bold">
                                                Drop images here
                                            </Heading>
                                            <Text fontWeight="light">or click to upload</Text>
                                        </Stack>
                                    </Stack>
                                </Box>
                                <Input
                                    type="file"
                                    height="100%"
                                    width="100%"
                                    position="absolute"
                                    top="0"
                                    left="0"
                                    opacity="0"
                                    aria-hidden="true"
                                    accept="image/*"
                                />
                            </Box>
                        </Box>
                    </AspectRatio>
                    <Input
                        m={4}
                        w={fields_width}
                        bgColor={'whitesmoke'}
                        type="text"
                        placeholder="Name of Activity"
                        onChange={(e) => setActivityname(e.target.value)}
                    ></Input>
                    <Checkbox>Is it Free?</Checkbox>
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
                            <Text fontSize='3xl'>Date and Time of Visit</Text>
                            <HStack spacing='24px'>
                                <Input size="md" type="datetime-local" onChange={(e) => setVisitDateStart(e.target.value)}/>
                                <Text fontSize='1xl'> to </Text>
                                <Input size="md" type="datetime-local" onChange={(e) => setVisitDateEnd(e.target.value)}/>
                            </HStack>

                        </Stack>
                    </Flex>
                    <Button w={fields_width} m={4} type="submit" onClick={onSubmit} bg="teal">
                        <font size={5} color={'white'}>Add to itinerary</font>
                    </Button>
                    <Spacer/>

                </VStack>
            </GridItem>
        </Grid>

    );
};

export default CreateItinerary;