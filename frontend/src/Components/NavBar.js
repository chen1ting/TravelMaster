import { useState } from "react";
import {
  Flex,
  Box,
  useColorModeValue,
  Menu,
  MenuButton,
  Button,
  Avatar,
  MenuList,
  MenuItem,
  Image,
} from "@chakra-ui/react";
import { ENDPOINT } from "../api/api";
import { useNavigate } from "react-router-dom";
import TMLogo from "../Components/TMLogo.png";
import { sendLogoutReq } from "../api/api";

const paddingSpace = "20px";

const NavButton = (displayString, navigateString) => {
  const navigate = useNavigate();
  return (
    <MenuButton
      px="2"
      py="1"
      rounded="md"
      onClick={() => {
        navigate(navigateString);
      }}
      _hover={{
        bg: useColorModeValue("blue.200", "blue.400"),
        fontWeight: "bold",
      }}
      m={2}
    >
      {displayString}
    </MenuButton>
  );
};

const Header = ({ imageUrl, setImageUrl }) => {
  const navigate = useNavigate();
  const [show] = useState(false);

  return (
    <Flex
      // mb={4}
      p={2}
      as="nav"
      alignItems="center"
      justify="space-between"
      w="100%"
      bg="cornflowerblue"
      h="11vh"
    >
      <Box
        w="200px"
        p="-1"
        paddingLeft={paddingSpace}
        cursor="pointer"
        onClick={() => navigate("/")}
      >
        <Image src={TMLogo} boxSize="24" />
      </Box>

      <Box
        display={{ base: show ? "block" : "none", md: "block" }}
        flexBasis={{ base: "100%", md: "auto" }}
        paddingRight={paddingSpace}
      >
        <Flex fontSize={"xl"} px="10">
          {/* Navbar buttons */}
          <Menu>
            {NavButton("Discover", "/discover")}
            {NavButton("Create an Itinerary", "/welcome")}
            {NavButton("Itineraries", "/itineraries")}
            {NavButton("Feedback", "/feedback")}
            {/*{NavButton('Bookings', '/bookings')}*/}
          </Menu>

          {/* User Icon (Profile) */}

          <Flex alignItems={"center"} px="3">
            <Menu>
              <MenuButton
                as={Button}
                rounded={"full"}
                variant={"link"}
                cursor={"pointer"}
                minW={0}
                paddingLeft={"14px"}
              >
                {/* Replace src with user image url, with the default as current image*/}
                <Avatar size={"md"} src={`${ENDPOINT}/avatars/` + imageUrl} />
              </MenuButton>

              {imageUrl ? (
                <MenuList justify="center">
                  <MenuItem
                    onClick={() => {
                      navigate("/profile");
                    }}
                    px="2"
                    py="1"
                    rounded="md"
                    _hover={{
                      color: "white",
                      bg: "blue.400",
                      fontWeight: "bold",
                    }}
                  >
                    Profile
                  </MenuItem>
                  <MenuItem
                    onClick={async () => {
                      await sendLogoutReq(
                        window.sessionStorage.getItem("session_token")
                      );
                      window.sessionStorage.clear();
                      setImageUrl("");
                      navigate("/");
                    }}
                    px="2"
                    py="1"
                    rounded="md"
                    _hover={{
                      color: "white",
                      bg: "blue.400",
                      fontWeight: "bold",
                    }}
                  >
                    Log Out
                  </MenuItem>
                </MenuList>
              ) : (
                <MenuList>
                  <MenuItem
                    onClick={() => {
                      navigate("/");
                    }}
                    px="2"
                    py="1"
                    rounded="md"
                    _hover={{
                      color: "white",
                      bg: "blue.400",
                      fontWeight: "bold",
                    }}
                  >
                    Sign In
                  </MenuItem>
                  <MenuItem
                    onClick={() => {
                      navigate("/Signup");
                    }}
                    px="2"
                    py="1"
                    rounded="md"
                    _hover={{
                      color: "white",
                      bg: "blue.400",
                      fontWeight: "bold",
                    }}
                  >
                    Create an account
                  </MenuItem>
                </MenuList>
              )}
            </Menu>
          </Flex>
        </Flex>
      </Box>
    </Flex>
  );
};

export default Header;
