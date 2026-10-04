;; Development manifest for vexilo.
;;
;; Most development in this repository is expected to happen through this
;; manifest. The repo uses Go 1.27.

(specifications->manifest
 (list "emacs"
       "gcc-toolchain"
       "git"
       "go@1.27"
       "go-golang-org-x-tools-godoc"
       "gopls"
       "podman"
       "podman-compose"
       "ripgrep"))