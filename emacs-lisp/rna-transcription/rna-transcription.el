;;; rna-transcription.el -- RNA Transcription (exercism)  -*- lexical-binding: t; -*-

;;; Commentary:

;;; Code:
(defun transcript (ch)
  "CH is the nucleotide."
  (pcase ch
    (?G ?C)
    (?C ?G)
    (?T ?A)
    (?A ?U)
    (_ (error "Not a NDA nuleotides"))))

(defun to-rna (strand)
  "Convert STRAND to rna."
  (apply #'string (mapcar #'transcript strand)))

;; (defconst dna-to-rna '((?G "C")
;;                        (?C "G")
;;                        (?T "A")
;;                        (?A "U")))
;; (defun to-rna (x)
;;   "Given DNA strand X, return the RNA complement."
;;   (mapconcat (lambda (elt)
;;                (or (car (alist-get elt dna-to-rna))
;;                    (error "%s is not  a valid dna nucleotide" elt))) x ""))

(provide 'rna-transcription)
;;; rna-transcription.el ends here
