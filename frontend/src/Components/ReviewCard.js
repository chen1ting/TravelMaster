import {Badge, Box, HStack, Image, Stack, Text} from "@chakra-ui/react";
import {StarIcon} from "@chakra-ui/icons";
let name= "Felix"
let date= "23 September 2022"
let rating = 4
let profileImg = "https://mdbootstrap.com/img/new/standard/city/042.webp"
let uploadImg1 = "https://mdbootstrap.com/img/new/standard/city/045.webp"
let uploadImg2 = "https://mdbootstrap.com/img/new/standard/city/046.webp"
let uploadImg3 = "https://mdbootstrap.com/img/new/standard/city/047.webp"
let reviewTitle = "Day ticket holder"
let reviewBody = "Brilliant Penguins"
let ReviewCard = ({}) => {

    return (
        <div className="app">
            <Box w="container.md" rounded="20px"
                 overflow="hidden" bg={"gray.200"} mt={10}>
                <HStack spacing='20px'>
                    <Image src= {profileImg}
                           alt="Card Image" boxSize="80px">
                    </Image>
                    <Text as="h2" fontWeight="normal" my={2}>
                        {name}
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
                                {date}
                            </Badge>
                        </Box>
                    </Stack>
                    <Stack align="center">
                        <Text as="h2" fontWeight="normal" my={2}>
                            {reviewTitle}
                        </Text>
                        <Text fontWeight="light">
                            {reviewBody}
                        </Text>
                    </Stack>
                </Box>
                <HStack spacing='10px'>
                    <Image src={uploadImg1}
                           alt="Card Image" boxSize="200px">
                    </Image>
                    <Image src={uploadImg2}
                           alt="Card Image" boxSize="200px">
                    </Image>
                    <Image src={uploadImg3}
                           alt="Card Image" boxSize="200px">
                    </Image>
                </HStack>
            </Box>
        </div>
    )
}

export default ReviewCard;