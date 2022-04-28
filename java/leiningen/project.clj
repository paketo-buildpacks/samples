(defproject demo "0.0.1-SNAPSHOT"
  :dependencies [[org.clojure/clojure "1.11.1"]
                 [io.pedestal/pedestal.service "0.5.10"]
                 [io.pedestal/pedestal.jetty "0.5.10"]
                 [ch.qos.logback/logback-classic "1.2.11" :exclusions [org.slf4j/slf4j-api]]
                 [org.slf4j/jul-to-slf4j "1.7.36"]
                 [org.slf4j/jcl-over-slf4j "1.7.36"]
                 [org.slf4j/log4j-over-slf4j "1.7.36"]]
  :resource-paths ["config"]
  :main ^:skip-aot lein-source.server
  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
