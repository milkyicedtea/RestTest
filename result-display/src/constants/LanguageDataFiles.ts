import cppData from "../../../wrk-results/cppcrow-app_benchmark_user-json.json";
import goData from "../../../wrk-results/gochi-app_benchmark_user-db-read.json";
import rustdata from "../../../wrk-results/rustaxum-app_benchmark_user-json.json";
import zigData from "../../../wrk-results/zigzap-app_benchmark_user-json.json";

import cppImage from "../../public/c++.png";
import goImage from "../../public/go.svg";
import rustImage from "../../public/rust.png";
import zigImage from "../../public/zig.svg";

// I haven't figured out why there are multiple Go files yet so I am leaving them out for now
export const LanguageDataFiles = [
  {
    lang_name: "C++",
    frameworks: ["crow"],
    image: cppImage,
    ...cppData,
  },
  {
    lang_name: "Go",
    frameworks: ["chi"],
    image: goImage,
    ...goData,
  },
  {
    lang_name: "Rust",
    frameworks: ["axum"],
    image: rustImage,
    ...rustdata,
  },
  {
    lang_name: "Zig",
    frameworks: ["zap"],
    image: zigImage,
    ...zigData,
  },
];
