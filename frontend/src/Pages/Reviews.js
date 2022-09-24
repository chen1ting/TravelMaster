import {useState} from 'react';
//import { useAuth } from '../lib/auth';
import {
    FormControl, FormLabel, Button, Input, Box, Flex, Grid, GridItem, Image, Badge, Text, Stack, Spacer, HStack
} from '@chakra-ui/react';
import {StarIcon} from "@chakra-ui/icons";

const fields_width = '52.5%';
let attractionName;
attractionName = "Penguin Feeding Show"
let showFree;
showFree = "Free"
let attrationDescription;
attrationDescription = "Have you met the Emperor Penguin?"
const rating = 4

const reviewList = [
    { name: "John", age: 24 },
    { name: "Linda", age: 19 },
    { name: "Josh", age: 33 }
];
const reviewCount = 3;
const renderReviews = () => {
    const result = [];
    for (let i = 0; i < reviewCount; i++) {
        
    }

    return <ul>{result}</ul>;
};

const Reviews = () => {

    //const { signIn } = useAuth();


    return (<Grid
        templateAreas={`"left_top right"
                            "left_bottom right"
                            `}
        gridTemplateRows={'50fr 50fr'}
        gridTemplateColumns={'1fr 2fr'}
        h='650px'
        // gap='1'
        // color='blackAlpha.700'
        fontWeight='bold'
    >
        <GridItem pl='2' bg='blue.200' area={'left_top'}>
            <h1>
                <font size={6}>
                    Reviews
                </font>
            </h1>

            <div className="app">
                <Box w="fit-content" rounded="20px" 
                    overflow="hidden" bg={"blue.500"} mt={10}>
                    <Image src="https://mdbootstrap.com/img/new/standard/city/041.webp"
                        alt="Card Image" boxSize="400px">
                    </Image>
                    <Box p={5}>
                    <Stack align="center">
                        <Badge variant="solid" colorScheme="green" 
                        rounded="full" px={2}>
                        {showFree}
                        
                        </Badge>
                    </Stack>
                    <Stack align="center">
                        <Text as="h2" fontWeight="normal" my={2} >
                        {attractionName}
                        </Text>
                        <Text fontWeight="light">
                        {attrationDescription}
                        </Text>
                    </Stack>
                    </Box>
                </Box>
            </div>
        </GridItem>

        <GridItem pl='2' bg='white' area={'left_bottom'}>
            
        </GridItem>

        <GridItem pl='2' bg='white' area={'right'}>
            <div className="app">
                <Box w="container.md" rounded="20px" 
                    overflow="hidden" bg={"gray.200"} mt={10}>
                    <HStack spacing='20px'>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/042.webp"
                            alt="Card Image" boxSize="80px">
                        </Image>
                        <Text as="h2" fontWeight="normal" my={2} >
                        Felix
                        </Text>
                    </HStack>
                    <Box p={5}>
                    <Stack align="center">
                    <Box display='flex' mt='2' alignItems='center'>
                        {Array(5)
                            .fill('')
                            .map((_, i) => (
                            <StarIcon
                                key={i}
                                color={i < rating ? 'teal.500' : 'gray.300'}
                            />
                            ))}
                        <Badge variant="solid" colorScheme="green" 
                        rounded="full" px={2}>
                        23 September 2022
                        </Badge>
                    </Box>
                    </Stack>
                    <Stack align="center">
                        <Text as="h2" fontWeight="normal" my={2} >
                        Day ticket holder
                        </Text>
                        <Text fontWeight="light">
                        Brilliant Penguins
                        </Text>
                    </Stack>
                    </Box>
                    <HStack spacing='10px'>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/045.webp"
                            alt="Card Image" boxSize="200px">
                        </Image>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/046.webp"
                            alt="Card Image" boxSize="200px">
                        </Image>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/047.webp"
                            alt="Card Image" boxSize="200px">
                        </Image>
                    </HStack>
                </Box>
            </div>
            <div className="app">
                <Box w="container.md" rounded="20px" 
                    overflow="hidden" bg={"gray.200"} mt={10}>
                    <HStack spacing='20px'>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/043.webp"
                            alt="Card Image" boxSize="80px">
                        </Image>
                        <Text as="h2" fontWeight="normal" my={2} >
                        Tom
                        </Text>
                    </HStack>
                    <Box p={5}>
                    <Stack align="center">
                    <Box display='flex' mt='2' alignItems='center'>
                        {Array(5)
                            .fill('')
                            .map((_, i) => (
                            <StarIcon
                                key={i}
                                color={i < rating ? 'teal.500' : 'gray.300'}
                            />
                            ))}
                        <Badge variant="solid" colorScheme="green" 
                        rounded="full" px={2}>
                        11 September 2022
                        </Badge>
                    </Box>
                    </Stack>
                    <Stack align="center">
                        <Text as="h2" fontWeight="normal" my={2} >
                        Annual Pass Visitor
                        </Text>
                        <Text fontWeight="light">
                        Penguins seem unhappy
                        </Text>
                    </Stack>
                    </Box>
                </Box>
            </div>
            <div className="app">
                <Box w="container.md" rounded="20px" 
                    overflow="hidden" bg={"gray.200"} mt={10}>
                    <HStack spacing='20px'>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/044.webp"
                            alt="Card Image" boxSize="80px">
                        </Image>
                        <Text as="h2" fontWeight="normal" my={2} >
                        Bob
                        </Text>
                    </HStack>
                    <Box p={5}>
                    <Stack align="center">
                    <Box display='flex' mt='2' alignItems='center'>
                        {Array(5)
                            .fill('')
                            .map((_, i) => (
                            <StarIcon
                                key={i}
                                color={i < rating ? 'teal.500' : 'gray.300'}
                            />
                            ))}
                        <Badge variant="solid" colorScheme="green" 
                        rounded="full" px={2}>
                        23 August 2022
                        </Badge>
                    </Box>
                    </Stack>
                    <Stack align="center">
                        <Text as="h2" fontWeight="normal" my={2} >
                        Foreign Tourist
                        </Text>
                        <Text fontWeight="light">
                        I have never seen penguins before in my life
                        </Text>
                    </Stack>
                    </Box>
                    <HStack spacing='10px'>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/048.webp"
                            alt="Card Image" boxSize="200px">
                        </Image>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/049.webp"
                            alt="Card Image" boxSize="200px">
                        </Image>
                        <Image src="https://mdbootstrap.com/img/new/standard/city/050.webp"
                            alt="Card Image" boxSize="200px">
                        </Image>
                    </HStack>
                </Box>
            </div>
            <Box h={20}/>
        </GridItem>
    </Grid>);
};

export default Reviews;