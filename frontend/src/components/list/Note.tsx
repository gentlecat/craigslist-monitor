import React from 'react';
import { css } from '@emotion/core';

type NoteProps = {
  note: string;
  onNoteUpdate: Function;
};

export const Note = ({ note, onNoteUpdate }: NoteProps) => {
  return (
    <textarea
      rows="3"
      cols="22"
      placeholder="Notes..."
      value={note}
      onChange={(event) => onNoteUpdate(event.target.value)}
      css={css`
        border: none;
        background-color: #e8e8e85e;
      `}
    />
  );
};

export default React.memo(Note);
