import { compose } from "@/shared/lib/compose";

import withToast from "./withToast";
import withQueryClient from "./withQueryClient";

/** TODO: add displaynames to HOCS */

export const withHocs = compose(withQueryClient, withToast);
