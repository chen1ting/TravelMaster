// common util functions
import { validateToken } from "../api/api";

// use this function in a useEffect hook to trigger redirect to homepage
// if a user has to be logged in first
const validSessionGuard = async (navigate) => {
  const ok = await validateToken(
    window.sessionStorage.getItem("session_token")
  );
  if (!ok) {
    // TODO: to improve this, clear invalid keys in session states
    // this is not necessary but makes sense to do so
    navigate("/");
  }
};

export { validSessionGuard };
