import { useState } from 'react';
import { Flex, Box, Text, Link, Menu, MenuButton, Button, Avatar, MenuList, MenuDivider} from '@chakra-ui/react';
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
      mb={4}
      p={2}
      as="nav"
      align="center"
      justify="space-between"
      wrap="wrap"
      w="100%"
      bg="blue.100"
    >
      <Flex
        wrap = "wrap"
        w = "20%"
        align = "flex-start"
        justify = "initial"
      >
        <Box w="200px" p = "-1">
          <Text as = "i" fontSize="5xl" fontWeight="bold">
            TM
          </Text>
        </Box>
        <Box w="200px">
          <Text fontSize="md" fontWeight="bold">
            Travel Master
          </Text>
        </Box>
      </Flex>
      <Box display={{ base: 'block', md: 'none' }} onClick={toggleMenu}>
        {show ? <CloseIcon /> : <HamburgerIcon />}
      </Box>

      <Box
        display={{ base: show ? 'block' : 'none', md: 'block' }}
        flexBasis={{ base: '100%', md: 'auto' }}
      >
        <Flex
        >
          <MenuItem to="/">Create an Itinerary</MenuItem>
          <MenuItem to="/itineraries">Itineraries</MenuItem>
          <MenuItem to="/bookings">Bookings</MenuItem>
          <Flex alignItems={'center'}>
            <Menu>
              <MenuButton
                as={Button}
                rounded={'full'}
                variant={'link'}
                cursor={'pointer'}
                minW={0}>
                <Avatar
                  size={'sm'}
                  src={
                    'https://www.cleanpng.com/png-person-silhouette-download-human-vector-1709279/://images.unsplash.com/photo-1493666438817-866a91353ca9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9'
                  }
                />
              </MenuButton>
              <MenuList justify = "center">
                <MenuItem>Log In</MenuItem>
                <MenuItem>Sign Up</MenuItem>
              </MenuList>
            </Menu>
          </Flex>
        </Flex>
      </Box>
    </Flex>
  );
};

export default Header;
