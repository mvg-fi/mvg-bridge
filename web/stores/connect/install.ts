import { MixinMessengerInstall, MixinMessengerName, UnisatInstall, UnisatName } from "~/helpers/constants";
import type { Wallet } from "~/types/wallet";

export const GetInstallLink = (w: Wallet): string => {
  switch (w.name) {
    case UnisatName:
      return UnisatInstall
    case MixinMessengerName:
      return MixinMessengerInstall
    // case FennecName:
      // break;
  }
  return ""
}