;;; maak.scm --- Maak build file for vexilo -*- mode: scheme; -*-
;;
;; SPDX-License-Identifier: AGPL-3.0-or-later or LGPL-3.0-or-later
;; Copyright © 2026 Urutau-Ltd <softwarelibre@urutau-ltd.org>
;;
;;   , _ ,      _    _            _                     _ _      _
;;  ( o o )    | |  | |          | |                   | | |    | |
;; /'` ' `'\   | |  | |_ __ _   _| |_ __ _ _   _ ______| | |_ __| |
;; |'''''''|   | |  | | '__| | | | __/ _` | | | |______| | __/ _` |
;; |\\'''//|   | |__| | |  | |_| | || (_| | |_| |      | | || (_| |
;;    """       \____/|_|   \__,_|\__\__,_|\__,_|      |_|\__\__,_|
;;
;; This program is free software: you can redistribute it and/or modify
;; it under the terms of the GNU General Public License as published by
;; the Free Software Foundation, either version 3 of the License, or (at
;; your option) any later version.
;;
;; This program is distributed in the hope that it will be useful, but
;; WITHOUT ANY WARRANTY; without even the implied warranty of
;; MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
;; General Public License for more details.
;;
;; You should have received a copy of the GNU General Public License
;; along with this program. If not, see <https://www.gnu.org/licenses/>.
;;
;; This file contains the vexilo build targets for maak
;; (codeberg.org/jjba23/maak).

(define-module (maak)
  #:declarative? #t
  #:use-module (maak maak))

(define %go-env "CGO_ENABLED=0")
(define %guix-shell "guix shell --network -m ./manifest.scm --")
(define %podman-compose "podman-compose --podman-path podman")

(define %version "1.2.1")       ; keep in sync with guix.scm
(define %tag (string-append "v" %version))
(define %remotes '("origin" "ro-mirror" "upstream"))

(define (test)
  "Run the Go test suite."
  ($ (list %go-env %guix-shell "go test -v ./...")))

(define (vet)
  "Run go vet."
  ($ (list %go-env %guix-shell "go vet ./...")))

(define (check)
  "Run test and vet."
  (test)
  (vet))

(define (env)
  "Open the Guix development shell."
  ($ (list "guix shell --network -m ./manifest.scm")))

(define (emacs)
  "Start Emacs in the Guix shell."
  ($ (list %guix-shell "emacs")))

(define (pkg)
  "Build the local Guix package definition."
  ($ (list "guix build -f ./guix.scm")))

(define (podman-build)
  "Build the CI container image."
  ($ (list %podman-compose "build ci")))

(define (podman-check)
  "Run the CI container."
  ($ (list %podman-compose "run --rm ci")))

(define (podman-shell)
  "Open a shell in the container."
  ($ (list %podman-compose "run --rm shell")))

(define (sync)
  "Push the main branch and tags to every remote."
  (for-each (lambda (remote)
              ($ (list "git push" remote "main" "--tags")))
            %remotes))

(define (registry)
  "Upload the tagged module to the Forgejo Go package registry."
  ($ (list "mkdir -p .cache"))
  ($ (list "git archive" "--format=zip"
           (string-append "--prefix=codeberg.org/urutau-ltd/vexilo@" %tag "/")
           "-o" ".cache/vexilo.zip" %tag))
  ($ (list "curl" "--fail"
           "--header" "\"Authorization: token $FORGEJO_TOKEN\""
           "--upload-file" ".cache/vexilo.zip"
           "https://sl.urutau-ltd.org/api/packages/urutau-ltd/go/upload")))

(define (release)
  "Check, tag, sync, then upload the release."
  (check)
  ($ (list "git tag" %tag))
  (sync)
  (registry))

(define (ci)
  "Run the full CI check."
  (check))

(define (default)
  "Run ci."
  (ci))