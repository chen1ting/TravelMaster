import { useState } from "react";
//import { useAuth } from '../lib/auth';
import { sendLoginReq } from "../api/api";

import {
  FormControl,
  FormLabel,
  Button,
  Input,
  Box,
  Flex,
  Grid,
  GridItem,
  Text,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

const fields_width = "52.5%";

const SignIn = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [notifMsg, setNotifMsg] = useState("");
  const [isError, setIsError] = useState(false);

  const navigate = useNavigate();

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

    await new Promise((resolve) => setTimeout(resolve, 1000)); // 1 sec
    // redirect to homepage
    navigate("/welcome");
  }

  return (
    <Grid
      templateAreas={`"left_top right"
                            "left_bottom right"
                            `}
      gridTemplateRows={"45fr 55fr"}
      gridTemplateColumns={"3fr 2fr"}
      h="100vh"
      // gap='1'
      // color='blackAlpha.700'
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
      <GridItem pl="2" bg="white" area={"left_bottom"}>
        {/*{To fill}*/}
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
                _hover={{ cursor: "pointer", bgColor: "green.100" }}
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
