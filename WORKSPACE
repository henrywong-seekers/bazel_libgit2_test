load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Rule repository
http_archive(
    name = "rules_foreign_cc",
    strip_prefix = "rules_foreign_cc-master",
    url = "https://github.com/bazelbuild/rules_foreign_cc/archive/master.zip",
)

load("@rules_foreign_cc//:workspace_definitions.bzl", "rules_foreign_cc_dependencies")

# Call this function from the WORKSPACE file to initialize rules_foreign_cc
#  dependencies and let neccesary code generation happen
#  (Code generation is needed to support different variants of the C++ Starlark API.).
#
#  Args:
#    native_tools_toolchains: pass the toolchains for toolchain types
#      '@rules_foreign_cc//tools/build_defs:cmake_toolchain' and
#      '@rules_foreign_cc//tools/build_defs:ninja_toolchain' with the needed platform constraints.
#      If you do not pass anything, registered default toolchains will be selected (see below).
#
#    register_default_tools: if True, the cmake and ninja toolchains, calling corresponding
#      preinstalled binaries by name (cmake, ninja) will be registered after
#      'native_tools_toolchains' without any platform constraints.
#      The default is True.
rules_foreign_cc_dependencies()

http_archive(
    name = "libgit2",
    strip_prefix = "libgit2-1.0.1",
    urls = ["https://github.com/libgit2/libgit2/releases/download/v1.0.1/libgit2-1.0.1.tar.gz"],
    build_file = "@//:libgit2.BUILD",
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8663604808d2738dc615a2c3eb70eba54a9a982089dd09f6ffe5d0e75771bc4f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.6/rules_go-v0.23.6.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.6/rules_go-v0.23.6.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

http_archive(
    name = "com_github_libgit2_git2go_v30",
    urls = ["https://github.com/libgit2/git2go/archive/v30.0.5.tar.gz"],
    strip_prefix = "git2go-30.0.5",
    build_file = "@//:git2go.BUILD",
)
