import {useState, useEffect} from "react";
import {sendLoginReq} from "../api/api";

import {
    FormControl,
    FormLabel,
    Button,
    Input,
    Box,
    Flex,
    Grid,
    GridItem,
    Text, Image, Center, Stack,
} from "@chakra-ui/react";
import {useNavigate} from "react-router-dom";
import {validSessionGuard} from "../common/common";
import Screenshot1 from "../LandingPageScreenshots/Screenshot1.png"
import Screenshot2 from "../LandingPageScreenshots/Screenshot2.png"
import Screenshot3 from "../LandingPageScreenshots/Screenshot3.png"
import Screenshot4 from "../LandingPageScreenshots/Screenshot4.png"

const fields_width = "52.5%";

const DisplayScreenshot = (imageSource, caption) => {
    return (
        <Stack>
            <Image src={imageSource} border={'1px'} borderColor='gray.500'/>
            <Center fontSize='xl'>{caption}</Center>
        </Stack>
    )
}

const SignIn = ({setImageUrl}) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [notifMsg, setNotifMsg] = useState("");
    const [isError, setIsError] = useState(false);

    const navigate = useNavigate();

    useEffect(() => {
        validSessionGuard(navigate, "", "/welcome");
    }, [navigate]);

    //const { signIn } = useAuth();

    async function onSubmit(e) {
        e.preventDefault();
        const data = await sendLoginReq(username, password);
        if (data == null) {
            setNotifMsg("invalid username or password");
            setIsError(true);
            return;
        }

        setNotifMsg(`Welcome back ${data.username}`);
        setIsError(false);
        // store data into session
        window.sessionStorage.setItem("uid", data.user_id);
        window.sessionStorage.setItem("username", data.username);
        window.sessionStorage.setItem("session_token", data.session_token);
        window.sessionStorage.setItem("avatar_file_name", data.avatar_file_name);
        setImageUrl(data.avatar_file_name);

        await new Promise((resolve) => setTimeout(resolve, 1000)); // 1 sec
        // redirect to homepage
        navigate("/welcome");
    }

    return (
        <Grid
            templateAreas={`"left_top left_top right"
                            "ss1 ss2 right"
                            "ss3 ss4 right"
                            `}
            gridTemplateRows={"30fr 40fr 40fr"}
            gridTemplateColumns={"32fr 32fr 36fr"}
            h="100vh"
            fontWeight="bold"
        >
            <GridItem pl="2" bg="blue.50" area={"left_top"}>
                <Box
                    position={"relative"}
                    top={"50%"}
                    left={"50%"}
                    transform={"translate(-50%,-50%)"}
                    textAlign={"center"}
                    color="black"
                >
                    <Text fontSize="4xl">Welcome to TravelMaster</Text>
                    <Text fontSize="2xl">
                        Generate new itineraries at your fingertips.
                    </Text>
                </Box>
            </GridItem>
            <GridItem p="3" bg="white" area={"ss1"}>
                {DisplayScreenshot(Screenshot1, "Generate itineraries for your visit!")}
            </GridItem>
            <GridItem p="3" bg="white" area={"ss2"}>
                {DisplayScreenshot(Screenshot2, "Discover activities, hidden gems and much more!")}
            </GridItem>
            <GridItem p="3" bg="white" area={"ss3"}>
                {DisplayScreenshot(Screenshot3, "Exciting activities awaits you!")}
            </GridItem>
            <GridItem p="3" bg="white" area={"ss4"}>
                {DisplayScreenshot(Screenshot4, "Edit your itinerary with a peace of mind!")}
            </GridItem>
            <GridItem pl="2" bg="blue.200" area={"right"}>
                <Box
                    position={"relative"}
                    top={"50%"}
                    left={"50%"}
                    transform={"translate(-50%,-50%)"}
                    textAlign={"center"}
                >
                    <FormControl b={"1px"} id="signin">
                        <Flex wrap="wrap" direction="column" align="center">
                            <FormLabel m={4}>
                                <font size={6}>Log in to TravelMaster</font>
                            </FormLabel>
                            {notifMsg && (
                                <Box
                                    px="6"
                                    py="2"
                                    borderRadius="13px"
                                    bgColor={isError ? "tomato" : "limegreen"}
                                >
                                    <Text fontSize="xl">{notifMsg}</Text>
                                </Box>
                            )}

                            <Box
                                w="100%"
                                display="flex"
                                alignItems="center"
                                justifyContent="center"
                            >
                                <Text fontSize="lg">Username</Text>
                                <Input
                                    color="black"
                                    m={4}
                                    w={fields_width}
                                    bgColor={"whitesmoke"}
                                    type="text"
                                    placeholder="Username"
                                    onChange={(e) => setUsername(e.target.value)}
                                ></Input>
                            </Box>
                            <Box
                                w="100%"
                                display="flex"
                                alignItems="center"
                                justifyContent="center"
                            >
                                <Text fontSize="lg">Password</Text>
                                <Input
                                    color="black"
                                    m={4}
                                    w={fields_width}
                                    bgColor={"whitesmoke"}
                                    type="password"
                                    placeholder="Password"
                                    onChange={(e) => setPassword(e.target.value)}
                                ></Input>
                            </Box>

                            <Button
                                w={fields_width}
                                m={4}
                                type="submit"
                                onClick={onSubmit}
                                bg="teal"
                            >
                                <font size={5} color={"white"}>
                                    Log In
                                </font>
                            </Button>
                            <Button
                                as="a"
                                onClick={() => {
                                    navigate("/Signup");
                                }}
                                _hover={{cursor: "pointer", bgColor: "green.100"}}
                                w={fields_width}
                                m={20}
                                type="submit"
                                bg="teal.50"
                            >
                                <font size={3} color={"teal"}>
                                    Don't have an account? Sign Up Here
                                </font>
                            </Button>
                        </Flex>
                    </FormControl>
                </Box>
            </GridItem>
        </Grid>
    );
};

export default SignIn;
