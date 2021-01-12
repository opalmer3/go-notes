//jshint esversion: 6

document.addEventListener("DOMContentLoaded", function(){

  document.querySelector('#add-form').addEventListener("submit", function(e){
    e.preventDefault();
    if(validateNoteInput(document.querySelector('.note-input').value)){
        this.submit();
    }
  });

  document.querySelectorAll('.delete-btn').forEach(item => {
      item.addEventListener('click', handleDelete);
  });

  document.querySelectorAll('.edit-btn').forEach(item => {
      item.addEventListener('click', handleEdit);
  });

});


function handleDelete(){
    const id = this.value;
    console.log(id);

    const form = document.createElement('form');
    const input = document.createElement('input');

    form.method = 'post';
    form.action = '/delete';

    input.name = 'Id';
    input.value = id;

    form.appendChild(input);

    document.body.appendChild(form);
    form.submit();
}


function handleEdit(){
  const id = this.value;
  const note = document.querySelector(`#note${id} .note-content`).textContent;
  document.querySelector(`#note${id} .note-btns`).style.display = 'none';

  const form = document.createElement('form');
  const input = document.createElement('textarea');
  const input1 = document.createElement('input');
  const btn = document.createElement('button');
  const btn1 = document.createElement('button');

  form.method = 'post';
  form.action = '/edit';
  form.classList.add('edit-form');

  input.name = 'note';
  input.classList.add('edit-note-input');
  input.value = note;

  input1.name = 'id';
  input1.setAttribute('type', 'hidden');
  input1.value = id;

  btn.type = 'submit';
  btn.classList.add('edit-note-btn');
  btn.textContent = 'Submit';

  btn1.type= 'button';
  btn1.classList.add('cancel-note-btn');
  btn1.textContent = 'Cancel';

  form.appendChild(input);
  form.appendChild(input1);
  form.appendChild(btn);
  form.appendChild(btn1);

  document.querySelector(`#note${id} .note-content`).textContent = "";
  document.querySelector(`#note${id} .note-content`).appendChild(form);

  document.querySelector(`#note${id} .edit-form`).addEventListener("submit", function(e){
    e.preventDefault();
    if(validateNoteInput(document.querySelector(`#note${id} .edit-note-input`).value)){
        this.submit();
    }
  });

  btn1.addEventListener('click', function(){
    document.querySelector(`#note${id} .edit-form`).remove();
    document.querySelector(`#note${id} .note-content`).textContent = note;
    document.querySelector(`#note${id} .note-btns`).style.display = 'block';
  });

  input.focus();
}


function validateNoteInput(input){
  if (input.trim() === '' ){ return false; }
  return true;
}
