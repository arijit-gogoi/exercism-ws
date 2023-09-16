;;; hamming.el --- Hamming (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:

;;; Code:
(require 'seq)

(defun hamming-distance (s1 s2)
  "Given two DNA strands (S1 and S2), output the Hamming distance."
  (if (= (length s1)
         (length s2))
      (seq-count #'not (seq-mapn #'= s1 s2))
    (error "DNA strands are of unequal lengths")))

(provide 'hamming)
;;; hamming.el ends here
