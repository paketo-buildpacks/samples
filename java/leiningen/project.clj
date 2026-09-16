(defproject demo "0.0.1-SNAPSHOT"
  :dependencies [[org.clojure/clojure "1.12.5"]
                 [io.pedestal/pedestal.service "0.7.2"]
                 [io.pedestal/pedestal.jetty "0.7.2"]
                 [ch.qos.logback/logback-classic "1.5.38" :exclusions [org.slf4j/slf4j-api]]
                 [org.slf4j/jul-to-slf4j "2.0.18"]
                 [org.slf4j/jcl-over-slf4j "2.0.18"]
                 [org.slf4j/log4j-over-slf4j "2.0.18"]]
  :resource-paths ["config"]
  :main ^:skip-aot lein-source.server
  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
