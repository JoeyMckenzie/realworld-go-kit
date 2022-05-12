load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "conduit-bin",
    embed = [":conduit-bin_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "conduit-bin_lib",
    srcs = ["main.go"],
    importpath = "github.com/joeymckenzie/realworld-go-kit/conduit-bin/conduit-bin",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_go_kit_log//level:go_default_library",
        "@com_github_joeymckenzie_realworld_go_kit_conduit_shared//config:go_default_library",
        "@com_github_joeymckenzie_realworld_go_kit_conduit_shared//persistence:go_default_library",
        "@com_github_joeymckenzie_realworld_go_kit_conduit_shared//utilities:go_default_library",
        "@io_entgo_ent//dialect/sql:go_default_library",
    ],
)
