// common util functions
import { validateToken } from "../api/api";
import { apiKey } from "../api/apiKey"

// use this function in a useEffect hook to trigger redirect to homepage
// if a user has to be logged in first
const validSessionGuard = async (navigate, pathIfNotOk, pathIfOk) => {
  const ok = await validateToken(
    window.sessionStorage.getItem("session_token")
  );
  if (!ok && pathIfNotOk !== "") {
    // TODO: to improve this, clear invalid keys in session states
    // this is not necessary but makes sense to do so
    navigate(pathIfNotOk);
  }
  if (ok && pathIfOk !== "") {
    navigate(pathIfOk);
  }
};

export { validSessionGuard, apiKey };
