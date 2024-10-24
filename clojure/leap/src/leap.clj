(ns leap)

(defn leap-year? [year] ;; <- argslist goes here
  (or (mod year 4) 
      (and (not (mod year 100)) (mod year 400))))
