import axios from "axios";
import { MixinAPIBaseURL } from "../constants";

export async function userMe(token: string) {
  let config = {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  };
  const result = await axios
    .get(MixinAPIBaseURL + "/me", config);
  return result.data ? result.data.data : {};
}