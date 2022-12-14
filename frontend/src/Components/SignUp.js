import {useState, useEffect} from "react";
import {validSessionGuard} from "../common/common";
import {sendSignupReq} from "../api/api";
import {
    FormControl,
    Button,
    Input,
    Box,
    Heading,
    FormLabel,
    Text,
} from "@chakra-ui/react";
import {useNavigate} from "react-router-dom";

const fields_width = {base: "250px", md: "500px"};
const SignUp = ({setImageUrl}) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [password2, setPassword2] = useState("");
    const [picture, setPicture] = useState(null);
    const [email, setEmail] = useState("");
    const [showError, setShowError] = useState(false);
    const [errMsg, setErrMsg] = useState("");
    //const { signIn } = useAuth();

    const navigate = useNavigate();

    useEffect(() => {
        validSessionGuard(navigate, "", "/welcome");
    }, [navigate]);

    async function onSubmit(e) {
        e.preventDefault();
        // might wanna consider adding a regex check for email format
        // and also password validation regex
        var bad =
            username === "" ||
            password === "" ||
            password !== password2 ||
            email === "" ||
            picture == null;
        setShowError(bad);
        if (bad) {
            if (username === "") {
                setErrMsg("A valid username is required.");
            } else if (password === "") {
                setErrMsg("A valid password is required.");
            } else if (password !== password2) {
                setErrMsg("Password confirmation does not match.");
            } else if (picture === null) {
                setErrMsg("An avatar is required.");
            } else {
                setErrMsg("A valid email is required.");
            }
            return;
        }
        setErrMsg(""); // always clear after

        const data = await sendSignupReq(username, password, email, picture);
        if (data && data.error != null) {
            setShowError(true);
            setErrMsg(data.error);
            return;
        }

        // store data into session
        window.sessionStorage.setItem("uid", data.user_id);
        window.sessionStorage.setItem("username", data.username);
        window.sessionStorage.setItem("session_token", data.session_token);
        window.sessionStorage.setItem("avatar_file_name", data.avatar_file_name);
        setImageUrl(data.avatar_file_name);

        // redirect to homepage
        navigate("/welcome");
    }

    return (
        <Box display="flex" justifyContent="center" h="94vh" bgColor="gray.50">
            <Box
                minW={{base: "500px", lg: "50%"}}
                bgColor="whitesmoke"
                display="flex"
                alignItems="center"
                flexDir="column"
                py="10"
            >
                <Heading>Sign up</Heading>
                {showError && errMsg !== "" && (
                    <Box
                        mt="5"
                        borderRadius="10px"
                        bgColor="red.500"
                        minW="50%"
                        textAlign="center"
                        p="2"
                    >
                        <Text>{errMsg}</Text>
                    </Box>
                )}
                <Box
                    mt="7"
                    rowGap="5px"
                    display="flex"
                    flexDir="column"
                    alignItems="center"
                >
                    <form>
                        <FormControl isRequired my="5">
                            <FormLabel>Username</FormLabel>
                            <Input
                                w={fields_width}
                                bgColor='gray.50'
                                borderColor='gray.500'
                                type="text"
                                placeholder="Enter your username"
                                onChange={(e) => setUsername(e.target.value)}
                            ></Input>
                        </FormControl>

                        <FormControl isRequired my="5">
                            <FormLabel>Email address</FormLabel>
                            <Input
                                w={fields_width}
                                bgColor='gray.50'
                                borderColor='gray.500'
                                type="text"
                                placeholder="Enter your email address"
                                onChange={(e) => setEmail(e.target.value)}
                            ></Input>
                        </FormControl>

                        <FormControl isRequired my="5">
                            <FormLabel>Password</FormLabel>
                            <Input
                                w={fields_width}
                                bgColor='gray.50'
                                borderColor='gray.500'
                                placeholder="Enter your password"
                                type="password"
                                onChange={(e) => setPassword(e.target.value)}
                            ></Input>
                        </FormControl>

                        <FormControl isRequired my="5">
                            <FormLabel>Confirm password</FormLabel>
                            <Input
                                w={fields_width}
                                bgColor='gray.50'
                                borderColor='gray.500'
                                type="password"
                                placeholder="Enter your password again"
                                onChange={(e) => setPassword2(e.target.value)}
                            ></Input>
                        </FormControl>

                        <FormControl isRequired my="5">
                            <FormLabel>Upload your avatar</FormLabel>
                            <input
                                type="file"
                                //style={{ display: 'none' }}
                                onChange={(e) => {
                                    setPicture(e.target.files[0]);
                                }}
                            />
                        </FormControl>

                        <Button
                            w={fields_width}
                            mt="5"
                            type="submit"
                            onClick={onSubmit}
                            bgColor="teal.500"
                            _hover={{bg: "teal.400"}}
                        >
                            Create Account
                        </Button>
                    </form>
                </Box>
            </Box>
        </Box>
    );
};

export default SignUp;
