import {useState} from 'react';
//import { useAuth } from '../lib/auth';
import {
    FormControl, FormLabel, Button, Flex, Grid, GridItem,
    AspectRatio,
    Box,
    BoxProps,
    Container,
    forwardRef,
    Heading,
    Input,
    Stack,
    Text,
    Checkbox, CheckboxGroup
} from "@chakra-ui/react";

import { motion, useAnimation } from "framer-motion";

//import UploadImage from './src/Components/UploadImage';
const fields_width = '52.5%';

const CreateEvent = () => {
    const [descriptionlocation, setDescriptionLocation] = useState('');
    const [addressevent, setAddressEvent] = useState('');
    const [descriptionevent, setDescriptionEvent] = useState('');

    const [eventname, setEventName] = useState('');

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
                    <Text fontSize='2xl'>Events</Text>
                </Box>

            </GridItem>



            <GridItem pl='2' bg='white' area={'left_bottom'}>
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
            </GridItem>



            <GridItem pl='2' bg='blue.200' area={'right'}>
                <Grid
                    templateAreas={`"top_left top_right"
                            "bottom_left bottom_right"
                            `}
                    gridTemplateRows={'50fr 50fr'}
                    gridTemplateColumns={'50fr 50fr'}
                    h='100vh'
                    // gap='1'
                    // color='blackAlpha.700'
                    fontWeight='bold'
                >
                    <GridItem pl='2' bg='blue.200' area={'top_left'}>
                        <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                             textAlign={"left"}>
                            <Flex
                                wrap="wrap"
                                direction="column"
                                align="center"
                            >
                                <Stack spacing={8}>
                                    <Text fontSize='3xl'>Description of Location</Text>
                                    <Text fontSize='3xl'>Address of Event</Text>
                                    <Text fontSize='3xl'>Description of Event</Text>
                                </Stack>
                            </Flex>
                        </Box>
                    </GridItem>


                    <GridItem pl='2' bg='blue.200' area={'top_right'}>
                        <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                             textAlign={"center"}>
                            <FormControl b={'1px'} id="CreateEvent1">
                                <Flex
                                    wrap="wrap"
                                    direction="column"
                                    align="left"
                                >
                                    <Input
                                        m={4}
                                        w={fields_width}
                                        bgColor={'whitesmoke'}
                                        type="text"
                                        placeholder=""
                                        onChange={(e) => setDescriptionLocation(e.target.value)}
                                    ></Input>
                                    <Input
                                        m={4}
                                        w={fields_width}
                                        bgColor={'whitesmoke'}
                                        type="text"
                                        placeholder=""
                                        onChange={(e) => setAddressEvent(e.target.value)}
                                    ></Input>
                                    <Input
                                        m={4}
                                        w={fields_width}
                                        bgColor={'whitesmoke'}
                                        type="text"
                                        placeholder=""
                                        onChange={(e) => setDescriptionEvent(e.target.value)}
                                    ></Input>
                                </Flex>
                            </FormControl>
                        </Box>
                    </GridItem>

                    <GridItem pl='2' bg='blue.200' area={'bottom_left'}>
                        <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                             textAlign={"center"}>
                            <Flex
                                wrap="wrap"
                                direction="column"
                                align="left"
                            >

                                <Stack spacing={4}>
                                    <Text fontSize='3xl'>Event Hours</Text>
                                    <Text fontSize='1xl'>Sunday</Text>
                                    <Text fontSize='1xl'>Monday</Text>
                                    <Text fontSize='1xl'>Tuesday</Text>
                                    <Text fontSize='1xl'>Wednesday</Text>
                                    <Text fontSize='1xl'>Thursday</Text>
                                    <Text fontSize='1xl'>Friday</Text>
                                    <Text fontSize='1xl'>Saturday</Text>
                                </Stack>
                            </Flex>
                        </Box>
                    </GridItem>

                    <GridItem pl='2' bg='blue.200' area={'bottom_right'}>
                    </GridItem>
                </Grid>
            </GridItem>
        </Grid>

    );
};

export default CreateEvent;