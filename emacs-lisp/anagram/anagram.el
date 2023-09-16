;;; anagram.el --- Anagram (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:

;;; Code:
(defun anagramsp (word1 word2)
  (let ((word1 (downcase word1))
        (word2 (downcase word2)))
    (and
     (not (equal word1 word2))
     (equal (seq-sort #'< word1) (seq-sort #'< word2)))))

(defun anagrams-for (word candidates)
  (seq-filter (lambda (candidate) (anagramsp word candidate)) candidates))

(provide 'anagram)
;;; anagram.el ends here
