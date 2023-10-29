import axios from "axios";
import { MixinAPIBaseURL } from "../constants";
import type { Asset } from "~/types/asset";

export async function networkAsset(a: Asset) {
  const result = await axios
    .get(MixinAPIBaseURL + "/network/assets/" + a.asset_id);
  return result.data ? result.data.data : {};
}