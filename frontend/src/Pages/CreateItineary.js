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

//import UploadImage from './src/Components/UploadImage';
const fields_width = '52.5%';

const CreateItinerary = () => {
    const [descriptionlocation, setDescriptionLocation] = useState('');
    const [addressevent, setAddressEvent] = useState('');
    const [descriptionevent, setDescriptionEvent] = useState('');
    const [eventname, setEventName] = useState('');
    const [visittimestart, setVisitTimeStart] = useState('');
    const [visittimeend, setVisitTimeEnd] = useState('');


    //const { signIn } = useAuth();

    function onSubmit(e) {
        e.preventDefault();

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
                    <Text fontSize='2xl'>Create Itineary</Text>
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
                        placeholder="Name of Event"
                        onChange={(e) => setEventName(e.target.value)}
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
                                <Text fontSize='xl'>Address of Event</Text>
                                <Input
                                    m={4}
                                    w={"96"}
                                    bgColor={'whitesmoke'}
                                    type="text"
                                    placeholder=""
                                    onChange={(e) => setAddressEvent(e.target.value)}
                                ></Input>
                            </HStack>
                            <HStack spacing='54px'>
                                <Text fontSize='xl'>Description of Event</Text>
                                <Input
                                    m={4}
                                    w={"96"}
                                    h={"56"}
                                    bgColor={'whitesmoke'}
                                    type="text"
                                    placeholder=""
                                    onChange={(e) => setDescriptionEvent(e.target.value)}
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
                            <Text fontSize='3xl'>Time of Visit</Text>
                            <HStack spacing='24px'>
                                <Input size="md" type="time" onChange={(e) => setVisitTimeStart(e.target.value)}/>
                                <Text fontSize='1xl'> to </Text>
                                <Input size="md" type="time" onChange={(e) => setVisitTimeEnd(e.target.value)}/>
                            </HStack>

                        </Stack>
                    </Flex>
                    <Button w={fields_width} m={4} type="submit" onClick={onSubmit} bg="teal">
                        <font size={5} color={'white'}>Add to itineary</font>
                    </Button>
                    <Spacer/>

                </VStack>
            </GridItem>
        </Grid>

    );
};

export default CreateItinerary;