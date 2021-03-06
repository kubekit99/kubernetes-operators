package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build", "docker_bundle")
load("@io_kubernetes_build//defs:build.bzl", "release_filegroup")

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//build/debs:all-srcs",
        "//build/release-tars:all-srcs",
        "//build/rpms:all-srcs",
        "//build/visible_to:all-srcs",
    ],
    tags = ["automanaged"],
)

docker_build(
    name = "busybox",
    debs = [
        "@busybox_deb//file",
    ],
    symlinks = {
        "/bin/sh": "/bin/busybox",
        "/usr/bin/busybox": "/bin/busybox",
        "/usr/sbin/busybox": "/bin/busybox",
        "/sbin/busybox": "/bin/busybox",
        # tee & mkfifo are used to send logs to /var/log and stdout, e.g. by kops
        "/bin/tee": "/bin/busybox",
        "/bin/mkfifo": "/bin/busybox",
        # Force creation of /tmp
        "/tmp/hostname": "/etc/hostname",
    },
)

# This list should roughly match kube::build::get_docker_wrapped_binaries()
# in build/common.sh.
DOCKERIZED_BINARIES = {
    "cloud-controller-manager": {
        "base": ":busybox",
        "target": "//cmd/cloud-controller-manager:cloud-controller-manager",
    },
    "kube-apiserver": {
        "base": ":busybox",
        "target": "//cmd/kube-apiserver:kube-apiserver",
    },
    "kube-controller-manager": {
        "base": ":busybox",
        "target": "//cmd/kube-controller-manager:kube-controller-manager",
    },
    "kube-scheduler": {
        "base": ":busybox",
        "target": "//plugin/cmd/kube-scheduler:kube-scheduler",
    },
    "kube-proxy": {
        "base": "@debian-iptables-amd64//image:image.tar",
        "target": "//cmd/kube-proxy:kube-proxy",
    },
}

[docker_build(
    name = binary + "-internal",
    base = meta["base"],
    cmd = ["/usr/bin/" + binary],
    debs = [
        "//build/debs:%s.deb" % binary,
    ],
    symlinks = {
        # Some cluster startup scripts expect to find the binaries in /usr/local/bin,
        # but the debs install the binaries into /usr/bin.
        "/usr/local/bin/" + binary: "/usr/bin/" + binary,
    },
) for binary, meta in DOCKERIZED_BINARIES.items()]

[docker_bundle(
    name = binary,
    images = {"gcr.io/google_containers/%s:{STABLE_DOCKER_TAG}" % binary: binary + "-internal"},
    stamp = True,
) for binary in DOCKERIZED_BINARIES.keys()]

[genrule(
    name = binary + "_docker_tag",
    srcs = [meta["target"]],
    outs = [binary + ".docker_tag"],
    cmd = "grep ^STABLE_DOCKER_TAG bazel-out/stable-status.txt | awk '{print $$2}' >$@",
    stamp = 1,
) for binary, meta in DOCKERIZED_BINARIES.items()]

genrule(
    name = "os_package_version",
    outs = ["version"],
    cmd = """
grep ^STABLE_BUILD_SCM_REVISION bazel-out/stable-status.txt \
    | awk '{print $$2}' \
    | sed -e 's/^v//' -e 's/[\+-]/_/g' \
    >$@
""",
    stamp = 1,
)

release_filegroup(
    name = "docker-artifacts",
    srcs = [":%s.tar" % binary for binary in DOCKERIZED_BINARIES.keys()] +
           [":%s.docker_tag" % binary for binary in DOCKERIZED_BINARIES.keys()],
)

# KUBE_CLIENT_TARGETS
release_filegroup(
    name = "client-targets",
    srcs = [
        "//cmd/kubectl",
        "//federation/cmd/kubefed",
    ],
)

# KUBE_NODE_TARGETS
release_filegroup(
    name = "node-targets",
    srcs = [
        "//cmd/kube-proxy",
        "//cmd/kubelet",
    ],
)

# KUBE_SERVER_TARGETS
# No need to duplicate CLIENT_TARGETS or NODE_TARGETS here,
# since we include them in the actual build rule.
release_filegroup(
    name = "server-targets",
    srcs = [
        "//cmd/cloud-controller-manager",
        "//cmd/hyperkube",
        "//cmd/kube-apiserver",
        "//cmd/kube-controller-manager",
        "//cmd/kubeadm",
        "//plugin/cmd/kube-scheduler",
        "//vendor/k8s.io/kube-aggregator",
    ],
)

# kube::golang::test_targets
filegroup(
    name = "test-targets",
    srcs = [
        "//cmd/gendocs",
        "//cmd/genkubedocs",
        "//cmd/genman",
        "//cmd/genswaggertypedocs",
        "//cmd/genyaml",
        "//cmd/kubemark",  # TODO: server platforms only
        "//cmd/linkcheck",
        "//cmd/mungedocs",
        "//federation/cmd/genfeddocs",
        "//test/e2e:e2e.test",
        "//test/e2e_node:e2e_node.test",  # TODO: server platforms only
        "//vendor/github.com/onsi/ginkgo/ginkgo",
    ],
)

# KUBE_TEST_PORTABLE
filegroup(
    name = "test-portable-targets",
    srcs = [
        "//federation/develop:all-srcs",
        "//hack:e2e.go",
        "//hack:federated-ginkgo-e2e.sh",
        "//hack:get-build.sh",
        "//hack:ginkgo-e2e.sh",
        "//hack/e2e-internal:all-srcs",
        "//hack/lib:all-srcs",
        "//test/e2e/testing-manifests:all-srcs",
        "//test/kubemark:all-srcs",
    ],
)
