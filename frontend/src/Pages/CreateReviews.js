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
    HStack,
    VStack,
    StackDivider,
    Stack,
    Text,
    Checkbox, CheckboxGroup, Spacer
} from "@chakra-ui/react";

import { motion, useAnimation } from "framer-motion";

//import UploadImage from './src/Components/UploadImage';
const fields_width = '52.5%';

const CreateReview = () => {
    const [rating, setRating] = useState('');
    const [uploadImg1, setuploadImg1] = useState('');
    const [uploadImg2, setUploadImg2] = useState('');
    const [uploadImg3, setUploadImg3] = useState('');
    const [reviewTitle, setReviewTitle] = useState('');
    const [reviewBody, setReviewBody] = useState('');
    let name= "Felix"
    let date= "23 September 2022"
    let profileImg = "https://mdbootstrap.com/img/new/standard/city/042.webp"
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
                                <Text fontSize='xl'>Review Title</Text>
                                <Input
                                    m={4}
                                    w={"96"}
                                    bgColor={'whitesmoke'}
                                    type="text"
                                    placeholder=""
                                    onChange={(e) => setReviewTitle(e.target.value)}
                                ></Input>
                            </HStack>
                        </Stack>

                    </Flex>
                    <Flex
                        wrap="wrap"
                        direction="column"
                        align="start"
                    >
                        <Stack spacing={8}>
                            <HStack spacing='24px'>
                                <Text fontSize='xl'>Review</Text>
                                <Input
                                    m={4}
                                    w={"96"}
                                    bgColor={'whitesmoke'}
                                    type="text"
                                    placeholder=""
                                    onChange={(e) => setReviewBody(e.target.value)}
                                ></Input>
                            </HStack>
                        </Stack>

                    </Flex>
                    <Button w={fields_width} m={4} type="submit" onClick={onSubmit} bg="teal">
                        <font size={5} color={'white'}>Publish Review</font>
                    </Button>
                    <Spacer/>

                </VStack>
            </GridItem>
        </Grid>

    );
};

export default CreateReview;