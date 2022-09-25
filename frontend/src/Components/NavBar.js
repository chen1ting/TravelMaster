import {useState} from 'react';
import {Flex, Box, Text, useColorModeValue, Link, Menu, MenuButton, Button, Avatar, MenuList} from '@chakra-ui/react';

const paddingSpace = '20px'
const textPadSpace = '7px'

const MenuItem = ({ children, to = '/'}) => {
    return (
        <Link
            px = "1"
            py = "1"
            rounded = "md"
            _hover = {{
                bg: useColorModeValue('blue.300', 'blue.400'),
                fontWeight: "bold",
            }}
            // change href to navigate
            href={to}>
            {children}
        </Link>
    );
};

const MenuTextAlign = inputString => {
    return (
        <Box position={'relative'}
             top={'50%'}
             left={'50%'}
             transform={'translate(-50%,-50%)'}
             textAlign={"left"}
             width='fit-content'
             paddingLeft={textPadSpace}
             paddingRight={textPadSpace}
        >
            {inputString}
        </Box>
    )
}

const Header = (props) => {

    const [show, setShow] = useState(false);
    return (
        <Flex
            // mb={4}
            p={2}
            as="nav"
            alignItems="center"
            justify="space-between"
            w="100%"
            bg="cornflowerblue"
            h="6vh"
        >
            <Box w="200px" p="-1" paddingLeft={paddingSpace}>
                <Text as="i" fontSize="5xl" fontWeight="bold">
                    TM
                </Text>
            </Box>

            <Box
                display={{base: show ? 'block' : 'none', md: 'block'}}
                flexBasis={{base: '100%', md: 'auto'}}
                paddingRight={paddingSpace}
            >
                <Flex fontSize={'xl'} px="10"
                >
                    <MenuItem to="/discover">
                        {MenuTextAlign("Discover")}
                    </MenuItem>
                    <MenuItem to="/createitinery">
                        {MenuTextAlign("Create an Itinerary")}
                    </MenuItem>
                    <MenuItem to="/itineraries">
                        {MenuTextAlign("Itineraries")}
                    </MenuItem>
                    <MenuItem to="/bookings">
                        {MenuTextAlign("Bookings")}
                    </MenuItem>
                    <Flex alignItems={'center'} px="3">
                        <Menu>
                            <MenuButton
                                as={Button}
                                rounded={'full'}
                                variant={'link'}
                                cursor={'pointer'}
                                minW={0}
                                paddingLeft={'14px'}
                            >
                                <Avatar
                                    size={'md'}
                                    src={
                                        'https://www.cleanpng.com/png-person-silhouette-download-human-vector-1709279/://images.unsplash.com/photo-1493666438817-866a91353ca9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=faces&fit=crop&h=200&w=200&s=b616b2c5b373a80ffc9636ba24f7a4a9'
                                    }
                                />
                            </MenuButton>
                            <MenuList justify="center">
                                <Box position={'relative'}
                                     top={'50%'}
                                     left={'50%'}
                                     transform={'translate(-50%,0)'}
                                     textAlign={"left"}
                                     width='fit-content'
                                >
                                    <MenuItem to="/profile">Profile</MenuItem>
                                    <MenuItem>Log Out</MenuItem>
                                </Box>
                            </MenuList>
                        </Menu>
                    </Flex>
                </Flex>
            </Box>
        </Flex>
    );


};

export default Header;
