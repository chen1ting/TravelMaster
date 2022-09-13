import { useState } from 'react';
//import { useAuth } from '../lib/auth';
import { FormControl, FormLabel, Button, Input, Box, Flex } from '@chakra-ui/react';
const fields_width = '52.5%';
const SignUp = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    //const { signIn } = useAuth();

    function onSubmit(e) {
        e.preventDefault();
        //signIn({ username, password });
    }

    return (
        <Box>
            <FormControl b={'1px'} id="signup">
                <Flex
                    wrap = "wrap"
                    direction = "column"
                    align = "center"
                >
                    <FormLabel m={4}>Sign Up</FormLabel>
                    <Input
                        m={4}
                        w = "40%"
                        type="text"
                        placeholder="Name"
                        onChange={(e) => setUsername(e.target.value)}
                    ></Input>
                    <Input
                        m={4}
                        w = "40%"
                        type="text"
                        placeholder="Username"
                        onChange={(e) => setUsername(e.target.value)}
                    ></Input>
                    <Input
                        m={4}
                        w = "40%"
                        type="text"
                        placeholder="Email Address"
                        onChange={(e) => setUsername(e.target.value)}
                    ></Input>
                    <Input
                        m={4}
                        w = "40%"
                        type="password"
                        placeholder="Password"
                        onChange={(e) => setPassword(e.target.value)}
                    ></Input>
                    <Input
                        m={4}
                        w = "40%"
                        type="password"
                        placeholder="Re-enter Password"
                        onChange={(e) => setPassword(e.target.value)}
                    ></Input>
                    <Button w={fields_width} m={4} type="submit" onClick={onSubmit} bg="teal">
                        <font size={5} color={'white'}>Create Account</font>

                    </Button>
                </Flex>

            </FormControl>
        </Box>
    );
};

export default SignUp;