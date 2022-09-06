import { useState } from 'react';
//import { useAuth } from '../lib/auth';
import { FormControl, FormLabel, Button, Input, Box, Flex } from '@chakra-ui/react';

const SignIn = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  //const { signIn } = useAuth();

  function onSubmit(e) {
    e.preventDefault();
    //signIn({ username, password });
  }

  return (
    <Box>
      <FormControl b={'1px'} id="signin">
        <Flex
        wrap = "wrap"
        direction = "column"
        align = "center"
        >
          <FormLabel m={4}>Sign In</FormLabel>
          <Input
            m={4}
            w = "40%"
            type="text"
            placeholder="username"
            onChange={(e) => setUsername(e.target.value)}
          ></Input>
          <Input
            m={4}
            w = "40%"
            type="password"
            placeholder="password"
            onChange={(e) => setPassword(e.target.value)}
          ></Input>
          <Button w={'40%'} m={4} type="submit" onClick={onSubmit} bg = "blue.75">
          Log In
          </Button>
        </Flex>
        
      </FormControl>
    </Box>
  );
};

export default SignIn;