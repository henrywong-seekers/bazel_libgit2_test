package(default_visibility = ["//visibility:public"])
load("@rules_foreign_cc//tools/build_defs:cmake.bzl", "cmake_external")

filegroup(
    name = "all",
    srcs = glob(["**"]),
)

cmake_external(
    name = "libgit2",
    lib_source = ":all",
    out_lib_dir = "lib64",
    shared_libraries = ["libgit2.so"],
)
