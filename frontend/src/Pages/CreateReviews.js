import {useState} from 'react';
import { sendCreateReviewReq } from "../api/apiCreateReviews";


import {
    Button, Flex, Grid, GridItem,
    Box,
    Input,
    HStack,
    VStack,
    StackDivider,
    Stack,
    Text,
    Spacer
} from "@chakra-ui/react";

import {useNavigate} from "react-router-dom";


//import UploadImage from './src/Components/UploadImage';
const fields_width = '52.5%';

const CreateReview = () => {
    const [rating, setRating] = useState('');
    const [uploadImg1, setUploadImg1] = useState('');
    const [uploadImg2, setUploadImg2] = useState('');
    const [uploadImg3, setUploadImg3] = useState('');
    const [reviewTitle, setReviewTitle] = useState('');
    const [reviewBody, setReviewBody] = useState('');

    const [showError, setShowError] = useState(false);
    const [errMsg, setErrMsg] = useState("");
    let name= "Felix"
    //let date= "23 September 2022"
    let profileImg = "https://mdbootstrap.com/img/new/standard/city/042.webp"
    const navigate = useNavigate();

    async function onSubmit(e) {
        e.preventDefault();
        // might wanna consider adding a regex check for email format
        // and also password validation regex
        var bad =
            reviewTitle === "" ||
            reviewBody === ""
        setShowError(bad);
        if (bad) {
            setErrMsg("A valid review and title is required.");
            return;
        }
        setErrMsg(""); // always clear after

        const data = await sendCreateReviewReq(name, new Date().toISOString().slice(0, 10), profileImg, reviewTitle, reviewBody, uploadImg1, uploadImg2, uploadImg3); //////TO CHANGE THE FUNCTION
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
                    <Text fontSize='2xl'>Events</Text>
                </Box>

            </GridItem>

            <GridItem pl='2' bg='blue.100' area={'left_bottom'}>
                <VStack
                    spacing={4}
                    align='center'>
                    <Spacer/>


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

                    <Flex
                        wrap="wrap"
                        direction="column"
                        align="start"
                    >
                        <Stack spacing={8}>
                            <h1>Upload Image</h1>
                            <HStack spacing='24px'>
                                <Input size="lg" type="file" onChange={(e) => setUploadImg1(e.target.value)}/>
                                <Input size="lg" type="file" onChange={(e) => setUploadImg2(e.target.value)}/>
                                <Input size="lg" type="file" onChange={(e) => setUploadImg3(e.target.value)}/>
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