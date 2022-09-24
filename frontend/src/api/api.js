import sha256 from "crypto-js/sha256";

// TODO: change to read from env
const ENDPOINT = "http://localhost:8080";

const sendSignupReq = async (user, pass, email) => {
  const rawResponse = await fetch(ENDPOINT + "/signup", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: user,
      hashed_password: sha256(pass).toString(),
      email: email,
    }),
  });

  if (rawResponse.status !== 201) {
    console.log("resp: " + rawResponse.status); // TODO: might wanna return an err message to display here
    return null;
  }
  const content = await rawResponse.json();
  return content;
};

// returns a boolean indicating if a token is still valid for the session
// TODO: implement this
const validateToken = (tokenStr) => {
  if (tokenStr === "" || tokenStr === null) {
    return false;
  }

  // TODO: call the backend /validate_token with the token
  // and parse resp

  return true;
};

export { sendSignupReq, validateToken };
