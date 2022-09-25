import {useState} from 'react';
//import { useAuth } from '../lib/auth';
import {
    FormControl, FormLabel, Button, Input, Box, Flex, Grid, GridItem, Text
} from '@chakra-ui/react';

const fields_width = '52.5%';

const SignIn = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    //const { signIn } = useAuth();

    function onSubmit(e) {
        e.preventDefault();
        //signIn({ username, password });

    }

    return (
        <Grid
            templateAreas={`"left_top right"
                            "left_bottom right"
                            `}
            gridTemplateRows={'45fr 55fr'}
            gridTemplateColumns={'3fr 2fr'}
            h='100vh'
            // gap='1'
            // color='blackAlpha.700'
            fontWeight='bold'
        >
            <GridItem pl='2' bg='blue.50' area={'left_top'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    <Text fontSize='4xl'>Welcome to TravelMaster</Text>
                    <Text fontSize='2xl'>Generate new itineraries at your fingertips.</Text>
                </Box>
                
            </GridItem>
            <GridItem pl='2' bg='white' area={'left_bottom'}>
                {/*{To fill}*/}
            </GridItem>
            <GridItem pl='2' bg='blue.200' area={'right'}>
                <Box position={'relative'} top={'50%'} left={'50%'} transform={'translate(-50%,-50%)'}
                     textAlign={"center"}>
                    <FormControl b={'1px'} id="signin">
                        <Flex
                            wrap="wrap"
                            direction="column"
                            align="center"
                        >
                            <FormLabel m={4}>
                                <font size={6}>
                                    Log in to TravelMaster
                                </font>
                            </FormLabel>
                            <Input
                                m={4}
                                w={fields_width}
                                bgColor={'whitesmoke'}
                                type="text"
                                placeholder="Username"
                                onChange={(e) => setUsername(e.target.value)}
                            ></Input>
                            <Input
                                m={4}
                                w={fields_width}
                                bgColor={'whitesmoke'}
                                type="password"
                                placeholder="Password"
                                onChange={(e) => setPassword(e.target.value)}
                            ></Input>
                            <Button w={fields_width} m={4} type="submit" onClick={onSubmit} bg="teal">
                                <font size={5} color={'white'}>Log In</font>
                            </Button>
                            <Button as="a" href= "SignUp" w={fields_width} m={20} type="submit" bg="teal.50">
                                <font size={3} color={'teal'}>Don't have an account? Sign Up Here</font>
                            </Button>

                        </Flex>
                    </FormControl>
                </Box>
            </GridItem>
        </Grid>
        
    );
};

export default SignIn;