import { useState } from 'react';
import { Flex, Box, Text, Link, useColorModeValue } from '@chakra-ui/react';
import { CloseIcon, HamburgerIcon } from '@chakra-ui/icons';

//Renders each item on the menu: Bookings, Itineraries, etc.
const MenuItem = ({ children, isLast, to = '/' }) => {
  return (
    <Text
      mb={{ base: isLast ? 0 : 8, sm: 0 }}
      mr={{ base: 0, sm: isLast ? 0 : 8 }}
      display="block"
    >
      <Link href={to}>{children}</Link>
    </Text>
  );
};

const Header = (props) => {
  const [show, setShow] = useState(false);
  const toggleMenu = () => setShow(!show);
  return (
    <Flex
      mb={8}
      p={8}
      as="nav"
      align="center"
      justify="space-between"
      wrap="wrap"
      w="100%"
      color = ""
    >
      <Box w="200px">
        <Text fontSize="lg" fontWeight="bold">
          Travel Master
        </Text>
      </Box>

      <Box display={{ base: 'block', md: 'none' }} onClick={toggleMenu}>
        {show ? <CloseIcon /> : <HamburgerIcon />}
      </Box>

      <Box
        display={{ base: show ? 'block' : 'none', md: 'block' }}
        flexBasis={{ base: '100%', md: 'auto' }}
      >
        <Flex
          align="center"
          justify={['center', 'space-between', 'flex-end', 'flex-end']}
          direction={['column', 'row', 'row', 'row']}
          pt={[4, 4, 0, 0]}
        >
          <MenuItem to="/">Create an Itinerary</MenuItem>
          <MenuItem to="/itineraries">Itineraries</MenuItem>
          <MenuItem to="/bookings">Bookings</MenuItem>
          <MenuItem to="/signin" isLast>
            Sign In
          </MenuItem>
        </Flex>
      </Box>
    </Flex>
  );
};

export default Header;
