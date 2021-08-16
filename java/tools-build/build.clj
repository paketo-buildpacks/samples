(ns build
    (:require [clojure.tools.build.api :as b]))

(def lib 'my/lib1)
(def version "1.2.0")
(def class-dir "target/classes")
(def basis (b/create-basis {:project "deps.edn"}))
(def uber-file (format "target/%s-%s-standalone.jar" (name lib) version))


(defn uber [_]
      ; Delete
      (b/delete {:path "target"})

      ; Prepare
      (b/write-pom {:class-dir class-dir
                    :lib lib
                    :version version
                    :basis basis
                    :src-dirs ["src"]})
      (b/copy-dir {:src-dirs ["src" "resources"]
                   :target-dir class-dir})

      ;Uber
      (b/compile-clj {:basis basis
                      :src-dirs ["src"]
                      :class-dir class-dir})
      (b/uber {:class-dir class-dir
               :uber-file uber-file
               :basis basis
               :main  "toolsbuild_source.server"}))
